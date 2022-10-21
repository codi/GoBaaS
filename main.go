package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	initDB()
	for i := 0; i < 10000; i++ {
		println(i)
		loadData()
	}
}

var client *mongo.Client

func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().
		ApplyURI("mongodb://mongodb.staryet.com").
		SetMaxPoolSize(1000)
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func loadData() {
	collection := client.Database("gobaas").Collection("app")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map(): %v\n", result.Map()["name"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
