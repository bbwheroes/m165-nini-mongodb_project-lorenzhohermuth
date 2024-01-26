package mongodb

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoPokemon struct {
    PokemonId		int32		   `bson:"id"`
    Name		string		   `bson:"name"`
    Weight		int32		   `bson:"weight"`
    Height		int32		   `bson:"height"`
    BaseExperience	int32		   `bson:"base_experience"`
    IsBaseFrom		bool		   `bson:"is_base_form"`
}

type MongoStat struct {
    Name		string		   `bson:"name"`
    TimeMs		int64		   `bson:"time_ms"`
    Points		int32		   `bson:"points"`
}

func (mp MongoPokemon) GetValue(value string) string {
    var out string = ""
    if value == "id" { out = strconv.Itoa(int(mp.PokemonId)) }
    if value == "name" { out = mp.Name }
    if value == "weight" { out = strconv.Itoa(int(mp.Weight / 10)) } // to get form hg to kg
    if value == "height" { out = strconv.Itoa(int(mp.Height * 10)) } // to get from dm to cm
    if value == "base_experience" { out = strconv.Itoa(int(mp.BaseExperience)) }
    if value == "is_base_form" {
	if mp.IsBaseFrom {
	    out = "is"
	}else {
	    out = "is not"
	}
    }
    return out
}

const db string = "webapp"

func GetExecutePokemon(bsonQuery bson.D) MongoPokemon{
    client, ctx, cancel := connect()
    collection := client.Database(db).Collection("pokemon")
    res := collection.FindOne(ctx, bsonQuery)
    var result MongoPokemon
    err := res.Decode(&result)
    if err != nil {
	log.Fatal(err)
    }
    defer deferFunc(client, ctx, cancel)
    return result
}

func GetExecuteStat(bsonQuery bson.D, limit int64) []MongoStat{
    client, ctx, cancel := connect()
    collection := client.Database(db).Collection("stats")
    opts := options.Find().SetSort(bsonQuery).SetLimit(limit)
    cur, err := collection.Find(ctx, bson.D{}, opts)
    var out []MongoStat
    if err != nil {
	log.Fatal(err)
    }
    for cur.Next(ctx) {
	var result MongoStat
	err = cur.Decode(&result)
	if err != nil {
	    log.Fatal(err)
	}
	out = append(out, result)
    }
    defer deferFunc(client, ctx, cancel)
    return out
}

func PutExecute(mp MongoStat) {
    client, ctx, cancel := connect()
    collection := client.Database(db).Collection("stats")
    _, err := collection.InsertOne(ctx, mp)
    if err != nil {
	log.Fatal(err)
    }
    defer deferFunc(client, ctx, cancel)
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
    return client, ctx, cancel
}

func deferFunc(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
    cancel()
    func() {
      if err := client.Disconnect(ctx); err != nil {
        panic(err)
	}
    }()
}
