package models

import (
	"context"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

var DB *mongo.Database

func ConnectDataBase() {

	clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017/")
	db, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	DB = db.Database("test")
	err = DB.Client().Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}
func GetSession() *mongo.Database {
	return DB

}
