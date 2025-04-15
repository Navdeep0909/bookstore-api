package user

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb://localhost:27017"
var DB = "BookStore"
var userCollection = "User"

func CreateMongoClient() *mongo.Client {
	fmt.Println("Inside the createMongoClient User")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil { 
		fmt.Println("Error connecting to MongoDB:", err)
    	return nil
	}
	return client
}

func GetCollection(client *mongo.Client, collection string) *mongo.Collection{
	fmt.Println("Inside the GetCollection User")
	return client.Database(DB).Collection(collection)
}

func InsertUser(collection string, data User) *mongo.InsertOneResult{
	fmt.Println("Inside the Insert User")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	userCollection := GetCollection(CreateMongoClient(), collection)
	fmt.Println("Inside the InsertUser method and printing the userCollection : ", userCollection)
	id, err := userCollection.InsertOne(ctx, data)
	if err!=nil{
		return nil
	}
	fmt.Println("Inside the Insert User after insertion")
	return id
}

func GetUserByEmail(collection string, filter interface{}) (User, error){
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	userCollection := GetCollection(CreateMongoClient(), collection)

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil{
		return user, err
	}
	return user, nil
}





