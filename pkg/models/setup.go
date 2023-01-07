package models

import (
	"context"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var DB *mongo.Database

func ConnectDataBase() {
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	err := godotenv.Load(".env")
	clientOptions := options.Client().ApplyURI("mongodb://" + DbHost + ":" + DbPort + "/")
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
