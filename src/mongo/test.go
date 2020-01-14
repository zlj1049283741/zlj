package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"reflect"
	"time"
)

type ActLog1 struct {
	//{“act”:”view”, “openid”:””, “unionid”: “”, “product_code”:””, “product_name”:””}
	Unionid     string `json:"unionid"`
	ActType     string `json:"action type"`
	ProductCode string `json:"code for product"`
}

func main() {
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

	collection := client.Database("panasonic_test").Collection("actlog")
	fmt.Println(collection)
	/**

	 */
	//{“act”:”view”, “openid”:””, “unionid”: “”, “product_code”:””, “product_name”:””}

	//random ten unionId
	uniId := make([]string, 0)
	for i := 0; i < 10; i++ {
		uniId = append(uniId, fmt.Sprintf("%d",  i))
	}
	fmt.Println(uniId)

	//actType
	types := [...]string{"view", "share"}
	fmt.Println(types)

	productCodes := make([]string, 0)
	for i := 0; i < 10; i++ {
		productCodes = append(productCodes, fmt.Sprintf("%s-%d", "prod", i))
	}

	fmt.Println(productCodes)

	for i := 0; i <= 1000; i++ {
		oneDoc := ActLog1{
			uniId[rand.Intn(10)],
			types[rand.Intn(2)],
			productCodes[rand.Intn(10)],
		}
		fmt.Println(oneDoc)
		fmt.Println("oneDoc TYPE:", reflect.TypeOf(oneDoc))
		result, insertErr := collection.InsertOne(ctx, oneDoc)
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
			panic(insertErr)
		} else {
			fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
			fmt.Println("InsertOne() API result:", result)

			// get the inserted ID string
			newID := result.InsertedID
			fmt.Println("InsertOne() newID:", newID)
			fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
		}

	}

	var result ActLog1
	// Get a MongoDB document using the FindOne() method
	err = collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		//fmt.Println("FindOne() ERROR:", err)
		panic(err)
	} else {
		//fmt.Println("FindOne() result:", result)
	}

	fmt.Println("--------------------------------------")
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip(0)
	findOptions.SetSort(bson.D{{"unionId", 1}})


	// Here's an array in which you can store the decoded documents
	var results []*ActLog1

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem ActLog1
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