/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"

	"github.com/ddr4869/go-microservices/src/movie-microservice/databases"
	"github.com/ddr4869/go-microservices/src/movie-microservice/models"
	"gopkg.in/mgo.v2/bson"
)

// Movie manages Movie CRUD
type Movie struct {
}

// COLLECTION of the database table
const (
	COLLECTION = "movies"
)

// GetAll gets the list of Movie
func (m *Movie) GetAll() ([]models.Movie, error) {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("movies")
	var movies []models.Movie
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	result := models.Movie{}
	for cur.Next(context.Background()) {
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		movies = append(movies, models.Movie{ID: result.ID, Name: result.Name, URL: result.URL, CoverImage: result.CoverImage, Description: result.Description})
	}
	defer cur.Close(context.Background())
	return movies, err
}

// GetByID finds a Movie by its id
func (m *Movie) GetByID(id string) (*models.Movie, error) {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("movies")
	var movie models.Movie
	cur, err := collection.Find(context.Background(), bson.M{"ID": id})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		err := cur.Decode(&movie)
		if err != nil {
			return nil, err
		}
	}
	defer cur.Close(context.Background())

	return &movie, err
}

// Insert adds a new Movie into database'
func (m *Movie) Insert(movie models.Movie) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("movies")
	_, err := collection.InsertOne(context.Background(), bson.M{"name": movie.Name, "url": movie.URL, "coverImage": movie.CoverImage, "description": movie.Description})
	if err != nil {
		return err
	}
	return nil
}

// Delete remove an existing Movie
func (m *Movie) Delete(movie models.Movie) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("movies")
	_, err := collection.DeleteOne(context.Background(), bson.M{"ID": movie.ID})
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing Movie
func (m *Movie) Update(movie models.Movie) error {
	collection := databases.Database.MgDbSession.Database("go-microservices").Collection("movies")
	_, err := collection.UpdateOne(context.Background(), bson.M{"ID": movie.ID}, bson.M{"$set": bson.M{"name": movie.Name, "url": movie.URL, "coverImage": movie.CoverImage, "description": movie.Description}})
	if err != nil {
		return err
	}
	return nil
}
