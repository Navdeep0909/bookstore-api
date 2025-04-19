package book

import (
	"context"
	"time"

	"github.com/navdeep0909/bookstore-api/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection = "Book"

func InsertBook(collection string, data Book) *mongo.InsertOneResult{
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	bookCollection := user.GetCollection(user.CreateMongoClient(), collection)
	id, err := bookCollection.InsertOne(ctx, data)
	if err!=nil{
		return nil
	}
	return id
}

func GetBooks(collection string, filter interface{}) ([]Book, error){
	var books []Book
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	bookCollection := user.GetCollection(user.CreateMongoClient(), collection)

	cursor, err := bookCollection.Find(ctx, filter)
	if err != nil{
		return nil, err
	}
	if err = cursor.All(context.TODO(), &books); err != nil{
		return nil, err
	}
	return books, nil
}

func GetBookById(collection string, filter interface{}) (Book, error){
	var book Book
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	bookCollection := user.GetCollection(user.CreateMongoClient(), collection)

	err := bookCollection.FindOne(ctx, filter).Decode(&book)
	if err != nil{
		return book, err
	}
	return book, nil
}

func UpdateBookInfo(collection string, filter interface{}, updateInfo interface{}) (*mongo.UpdateResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	bookCollection := user.GetCollection(user.CreateMongoClient(), collection)

	result, err := bookCollection.UpdateOne(ctx, filter, updateInfo)
	if err != nil{
		return  nil, err
	}
	return result, nil
}

func DeleteBookById(collection string, filter interface{}) *mongo.DeleteResult{
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	bookCollection := user.GetCollection(user.CreateMongoClient(), collection)

	result, err := bookCollection.DeleteOne(ctx, filter)
	if err != nil{
		return nil
	}
	return result
}