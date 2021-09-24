package storage

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Context context.Context
	Client  *mongo.Client
	*mongo.Database
}

type ShortUrl struct {
	Id          string `bson:"_id,required"`
	OriginalUrl string `bson:"original_url,required"`
	Visits      uint32 `bson:"visits"`
}

func NewDatabase(context context.Context, uri string, name string) *Database {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context, clientOptions)
	if err != nil {
		log.Panicln(err)
	}
	database := client.Database(name)

	return &Database{context, client, database}
}

func (db *Database) Insert(document interface{}, collectionName string) (id interface{}, err error) {
	collection := db.Collection(collectionName)
	res, err := collection.InsertOne(db.Context, document)
	if err != nil {
		return nil, fmt.Errorf("cannot insert document to database: %w", err)
	}
	id = res.InsertedID
	return
}

func (db *Database) FindOne(id interface{}, collectionName string) (res *mongo.SingleResult, err error) {
	collection := db.Collection(collectionName)
	res = collection.FindOne(db.Context, bson.M{"_id": id})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	return
}

func (db *Database) UpdateVisits(id interface{}, visits uint32, collectionName string) (res *mongo.UpdateResult, err error) {
	collection := db.Collection(collectionName)
	res, err = collection.UpdateOne(
		db.Context,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"visits", visits}}},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("cannot update visits: %w", err)
	}

	return
}

func (db *Database) FindShortUrl(id interface{}, collectionName string) (document *ShortUrl, err error) {
	document = &ShortUrl{}
	res, err := db.FindOne(id, collectionName)
	if err != nil {
		return nil, err
	}
	err = res.Decode(document)
	if err != nil {
		return nil, fmt.Errorf("error during decoding document: %w", err)
	}
	return
}
