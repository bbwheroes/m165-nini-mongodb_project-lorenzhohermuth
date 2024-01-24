package mongodb

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoPokemon struct {
    ID			primitive.ObjectID `bson:"_id"`
    PokemonId		int32		   `bson:"id"`
    Name		string		   `bson:"name"`
    Weight		int32		   `bson:"weight"`
    Height		int32		   `bson:"height"`
    BaseExperience	int32		   `bson:"base_experience"`
    IsBaseFrom		bool		   `bson:"is_base_form"`
}

func (mp MongoPokemon) GetValue(value string) string {
    var out string = ""
    if value == "id" { out = strconv.Itoa(int(mp.PokemonId)) }
    if value == "name" { out = mp.Name }
    if value == "weight" { out = strconv.Itoa(int(mp.Weight)) }
    if value == "height" { out = strconv.Itoa(int(mp.Height)) }
    if value == "base_experience" { out = string(mp.BaseExperience) }
    if value == "is_base_form" {
	if mp.IsBaseFrom {
	    out = "true"
	}else {
	    out = "false"
	}
    }
    return out
}

const db string = "webapp"
const coll string = "pokemon"

func Execute(bsonQuery bson.D) MongoPokemon{
    client, ctx, cancel := connect()
    collection := client.Database(db).Collection(coll)
    res := collection.FindOne(ctx, bsonQuery)
    var result MongoPokemon
    err := res.Decode(&result)
    if err != nil {
	log.Fatal(err)
    }
    defer deferFunc(client, ctx, cancel)
    return result
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
