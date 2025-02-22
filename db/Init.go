package db

import (
	"context"
	"log"
	"sma/dto"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL     = "mongodb://localhost:27017"
	dataBaseName = "analytics"
)

func Init() {
	if DB != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, error := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if error != nil {
		log.Fatalln("Mongo DB Connect Error : ", error)
	}
	log.Println("mongoDB Connect Succesfuly")
	db := client.Database(dataBaseName)
	DB = &analyticsDB{client: client, dataBase: db}
	adminUserInsert()
}

func adminUserInsert() {
	user := &dto.User{
		UserName:    "admin",
		Password:    "12345678",
		CustomerID:  "111111111111",
		ServiceType: "PREMIUM",
		Name:        "Admin",
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	DB.InsertOne(user)
}
