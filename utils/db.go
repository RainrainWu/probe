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

func init() {
	connectDB()
}

func connectDB() {
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Logger.Error("Failed to connect with mongodb")
		return
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Error("No response heard from mongodb")
		return
	}
	ReportColl = Client.Database("test").Collection("rain")
}

func WriteReport(doc interface{}) {
	
	_, err := ReportColl.InsertOne(context.TODO(), doc)
	if err != nil {
		Logger.Warn("Failed to write report")
	}
}

func ReadReportRaw(id string) Report {

	var result Report
	filter := bson.D{{"index", id}}
	err := ReportColl.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		Logger.Warn("Failed to read report")
	}
	return result
}

func ReadReport(id string) string {

	var result Report
	filter := bson.D{{"index", id}}
	err := ReportColl.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		Logger.Warn("Failed to read report")
	}
	data, _ := json.MarshalIndent(result, "", "  ")
	return string(data)
}
