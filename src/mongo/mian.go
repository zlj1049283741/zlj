package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type ActLog struct {
	//act_log_YM（id, uid, tag, act_type, target_value, duration, create_time)
	Uid string `json:"unionid"`
	Tag string `json:"tag"`
	ActType string `json:"action type"`
	TargetValue string `json:"target_value"`
	Duration string `json:"duration"`
	CreateTime string `json:"create_time"`
}

func ShowActLog(uid string) {
	/**
	  import test data to mongo db
	  db:test/user_
	*/

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://test:123456@localhost:27017/panasonic_test?authSource=test&authMechanism=SCRAM-SHA-1"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	//db.createUser({ user:'test',pwd:'123456',roles:[ { role:'readWrite', db: 'panasonic_test'}]});

	collection := client.Database("panasonic_test").Collection("act_log_YM")
	fmt.Println(collection)
	/**

	 */
	//act_log_YM（id, uid, tag, act_type, target_value, duration, create_time)
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip(2)//分页
	findOptions.SetSort(bson.D{{"uid", -1}})//排序

	var results []*ActLog

	cur, err := collection.Find(context.TODO(),bson.M{"uid":uid} , findOptions)//bson.bson.D{{}}
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem ActLog
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
		fmt.Println(elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

}

func main()  {
	ShowActLog("uni-7")
}