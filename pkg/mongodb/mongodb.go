package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type inputFunc func()

const db string = "webapp"
const coll string = "pokemon"

func execute(bsonQuery bson.D, bsonProj bson.D) {
    client, ctx, cancel := connect()
    collection := client.Database(db).Collection(coll)
    opts := options.Find().SetProjection(bsonProj)
    cur, err := collection.Find(ctx, bsonQuery, opts)
    if err != nil { log.Fatal(err)}
    for cur.Next(ctx) {
	var result bson.D
	err := cur.Decode(&result)
	if err != nil {
	    log.Fatal(err)
	}
	log.Println(result)
    }
    fmt.Println("Executed")
    defer deferFunc(cur,client, ctx, cancel)
}

func connect() (*mongo.Client, context.Context, context.CancelFunc){
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    clientOpts := options.Client().ApplyURI("mongodb://mongo:secretmongo@localhost:27017")
    client, connectErr := mongo.Connect(ctx, clientOpts)
    if connectErr != nil {
	panic(connectErr)
    }
    if pingErr := client.Ping(ctx, readpref.Primary()) ; pingErr != nil {
	panic(pingErr)
    }
    fmt.Println("Connected")
    return client, ctx, cancel
}

func deferFunc(cur *mongo.Cursor, client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
    cur.Close(ctx)
    cancel()
    func() {
      if err := client.Disconnect(ctx); err != nil {
        panic(err)
	}
    }()
    fmt.Println("Disconnected")
}
