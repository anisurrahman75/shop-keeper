package config

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func DBConnect() (*firestore.Client, error) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("/home/anisur/Downloads/stock-management-435-firebase-adminsdk-qwnic-b3bc661be3.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return client, nil
}
