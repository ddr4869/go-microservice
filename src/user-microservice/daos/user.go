/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"

	"github.com/ddr4869/go-microservices/src/user-microservice/databases"
	"github.com/ddr4869/go-microservices/src/user-microservice/models"
	"github.com/ddr4869/go-microservices/src/user-microservice/utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	var users []models.User
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	result := models.User{}
	for cur.Next(context.Background()) {
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		users = append(users, models.User{ID: result.ID, Name: result.Name, Password: result.Password})
	}
	defer cur.Close(context.Background())
	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (*models.User, error) {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	var user models.User
	cur, err := collection.Find(context.Background(), bson.M{"ID": id})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
	}
	defer cur.Close(context.Background())

	return &user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{"ID": id})
	if err != nil {
		return err
	}
	return nil
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"name": name, "password": password}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	user.ID = bson.NewObjectId()
	_, err := collection.InsertOne(context.Background(), bson.M{"ID": user.ID, "name": user.Name, "password": user.Password})
	return err
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{"ID": user.ID})
	if err != nil {
		return err
	}
	return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("users")
	_, err := collection.UpdateOne(context.Background(), bson.M{"ID": user.ID}, bson.M{"$set": bson.M{"name": user.Name, "password": user.Password}})
	if err != nil {
		return err
	}
	return nil
}
