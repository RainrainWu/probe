package main

import (
	"fmt"
	"log"
	"context"

	"github.com/RainrainWu/probe/pkg/utils"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client 		*mongo.Client
	ReportColl	*mongo.Collection
)

type Trainer struct {
    Name string
    Age  int
    City string
}

func Connect

func main() {
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.HandleErr(err, "Failed to connect with mongodb")
	
	err = client.Ping(context.TODO(), nil)
	utils.HandleErr(err, "No response heard from mongodb")

	ReportColl := client.Database("test").Collection("rain")
}

func CreateReport(doc interface{}) {
	
	result, err := ReportColl.InsertOne(context.TODO(), doc)
	HandleErr(err)
	fmt.Println("Inserted doc: ", result)
}
