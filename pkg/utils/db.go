package utils

import (
	"context"
	"encoding/json"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client 		*mongo.Client
	ReportColl	*mongo.Collection
)

func ConnectDB() {
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err := mongo.Connect(context.TODO(), clientOptions)
	HandleErr(err, "Failed to connect with mongodb")

	err = Client.Ping(context.TODO(), nil)
	HandleErr(err, "No response heard from mongodb")

	ReportColl = Client.Database("test").Collection("rain")
}

func WriteReport(doc interface{}) {
	
	_, err := ReportColl.InsertOne(context.TODO(), doc)
	HandleErr(err, "Failed to write report")
}

func ReadReport(id string) string {

	var result Report
	filter := bson.D{{"index", id}}
	err := ReportColl.FindOne(context.TODO(), filter).Decode(&result)
	HandleErr(err, "Failed to read report")
	
	data, _ := json.MarshalIndent(result, "", "  ")
	return string(data)
}
