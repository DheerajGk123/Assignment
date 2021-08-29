package Database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to database
func DBconnect() (*mongo.Client, error) {
	fmt.Println("Executing DBconnect()")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions) // Connect to the database
	if err != nil {
		log.Panic("Databae connection Failed : ", err)
		return client, errors.New("Database Connction Failed")
	}
	return client, nil
}

// disconnect from database
func DBdisconnect(client *mongo.Client) error {
	fmt.Println("Executing DBdisconnect()")
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Panic("Database connection Failed : ", err)
		return errors.New("Failed to Disconnect")
	}
	return nil
}

// Insert data to database
func InsertData(client *mongo.Client, u1 interface{}) error {
	fmt.Println("Executing InsertData()")
	_, e := client.Database("Assignment").Collection("User").InsertOne(context.TODO(), u1) // Insert data

	if e != nil {
		return errors.New("Database Insertion Failed")
	}
	return nil
}

// Retrive Database Data
func RetriveData(client *mongo.Client, id string, Robj interface{}) error {
	fmt.Println("Executing RetriveData()")
	fmt.Println("id is :", id)
	filter := bson.D{{"id", id}}
	err := client.Database("Assignment").Collection("User").FindOne(context.TODO(), filter).Decode(Robj)
	if err != nil {
		return errors.New("Failed to retrive data from database")
	}
	return nil
}

// Update Database Data
func UpdateDBData(client *mongo.Client, filter primitive.D, update primitive.D) error {
	_, e := client.Database("Assignment").Collection("User").UpdateOne(context.TODO(), filter, update)
	if e != nil {
		return errors.New("Failed to Update Database")
	}
	return nil
}

// Delete Database Data
func DeleteDBData(client *mongo.Client, id string) error {
	fmt.Println("Executing DeleteData()")
	fmt.Println("id is :", id)
	filter := bson.D{{"id", id}}
	_, err := client.Database("Assignment").Collection("User").DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Failed to retrive data from database")
	}
	return nil
}
