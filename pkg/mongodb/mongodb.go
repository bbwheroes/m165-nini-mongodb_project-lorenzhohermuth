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

var mongoContext *context.Context
var mongoClient *mongo.Client

func connect(fn inputFunc) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    clientOpts := options.Client().ApplyURI("mongodb://mongo:secretmongo@localhost:27017")
    client, err := mongo.Connect(ctx, clientOpts)
    err = client.Ping(ctx, readpref.Primary())
    mongoClient = client 
    mongoContext = &ctx
    fmt.Println("Connected")
    fn()
    fmt.Println("Executed")
    defer cancel()
    defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
	}
    }()
    defer fmt.Println("Disconnected")
}

func creatGetFunc(db string, coll string) func() {
    return func() {
	collection := mongoClient.Database(db).Collection(coll)
	cur, err := collection.Find(*mongoContext, bson.D{})
	if err != nil { log.Fatal(err)}
	for cur.Next(*mongoContext) {
	    var result bson.D
	    err := cur.Decode(&result)
	    if err != nil {
		log.Fatal(err)
	    }
	    log.Println(result)
	}
	defer cur.Close(*mongoContext)
    }
}
