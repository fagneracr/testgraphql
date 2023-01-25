package models

import (
	"apitest/pkg/token"
	"context"
	"html"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username"  bson:"username"`
	Password string             `json:"password" bson:"password"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, DB *mongo.Database) (string, error) {

	var err error

	u := User{}
	result := DB.Collection("users").FindOne(context.TODO(), bson.M{"username": username}, nil)
	err = result.Decode(&u)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser(DB *mongo.Database) (*User, error) {
	// user := User{
	// 	Username: u.Username, Password: u.Password,
	// }
	u.BeforeSave()
	if DB == nil {
		return nil, errors.New("Error on DB")
	}
	_, err := DB.Collection("users").InsertOne(ctx, &u, nil)
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func GetUserByID(uid primitive.ObjectID, DB *mongo.Database) (User, error) {

	var u User
	if err := DB.Collection("users").FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&u); err != nil {
		return u, errors.New("User not found")
	}

	//u.PrepareGive()

	return u, nil

}
