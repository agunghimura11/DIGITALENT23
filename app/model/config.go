package model

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
)

var client *db.Client
var ctx context.Context

func init(){
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://dts23-aae6f.firebaseio.com/",
	}

	opt := option.WithCredentialsFile("firebase-admin-sdk.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing databse client:", err)
	}
}