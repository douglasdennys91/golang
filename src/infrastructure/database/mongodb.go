package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type MongoDB struct {}

type IMongoDB interface {
	ConnectMongoDB() (*mongo.Database, error)
	GetAll(table string) ([]bson.M, error)
	Save(table string, data interface{}) (primitive.ObjectID, error)
	GetByID(table string, id primitive.ObjectID) (bson.M, error)
	GetParam(table string, params interface{}) (bson.M, error)
	Update(table string, id primitive.ObjectID, data interface{}) (interface{}, error)
	Delete(table string, id primitive.ObjectID) (bool, error)
}


func (m *MongoDB) ConnectMongoDB() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
	if err != nil {
		return nil, err
	}
	db := client.Database(os.Getenv("DB_NAME"))
	return db, nil
}

func (m *MongoDB) GetAll(table string) ([]bson.M, error) {
	collection, err := m.ConnectMongoDB()
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Collection(table).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var elements []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var element bson.M
		err := cur.Decode(&element)
		if err != nil {
			return nil, err
		}

		elements = append(elements, element)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return elements, nil
}

func (m *MongoDB) Save(table string, data interface{}) (interface{}, error) {
	collection, err := m.ConnectMongoDB()
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Collection(table).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return cur.InsertedID, nil
}

func (m *MongoDB) GetByID(table string, id primitive.ObjectID) (bson.M, error) {
	collection, err := m.ConnectMongoDB()
	if err != nil {
		return nil, err
	}

	var element bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.Collection(table).FindOne(ctx, bson.M{"_id": id}).Decode(&element)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func (m *MongoDB) GetParam(table string, params interface{}) (bson.M, error) {
	collection, err := m.ConnectMongoDB()
	if err != nil {
		return nil, err
	}

	var element bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.Collection(table).FindOne(ctx, params).Decode(&element)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func (m *MongoDB) Update(table string, id primitive.ObjectID, data interface{}) (bson.M, error) {
	collection, err := m.ConnectMongoDB()
	if err != nil {
		return nil, err
	}

	var element bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.Collection(table).FindOneAndUpdate(ctx, bson.M{"_id": id}, data).Decode(&element)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func (m *MongoDB) Delete(table string, id primitive.ObjectID) (bool, error) {
	collection, err := m.ConnectMongoDB()
	if err != nil {
		return false, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = collection.Collection(table).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return true, nil
}
