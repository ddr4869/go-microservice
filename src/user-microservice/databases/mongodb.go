/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package databases

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbSession  *mongo.Client
	Databasename string
}

var ctx = context.Background()

// Init initializes mongo database
func (db *MongoDB) Init() error {
	var err error
	db.MgDbSession, err = mongo.Connect(ctx)
	if err != nil {
		fmt.Println("Can't connect to mongo, go error: ", err)
		return err
	}
	return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error {
	var err error
	collection := db.MgDbSession.Database("go-microservices").Collection("users")
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Can't count user, go error: ", err)
		return err
	}
	if count == 0 {
		_, err = collection.InsertOne(context.Background(), bson.M{"name": "admin", "password": "Admin@123"})
		if err != nil {
			fmt.Println("Can't insert user, go error: ", err)
			return err
		}
	}
	return nil
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Disconnect(ctx)
	}
}
