package middleware

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	firebaseApp  *firebase.App
	firebaseAuth *auth.Client
)

func init() {
	ctx := context.Background()
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credentialsPath == "" {
		if os.Getenv("ENVIRONMENT") == "production" {
			credentialsPath = "/app/credentials/firebase.json"
		} else {
			credentialsPath = "./credentials/firebase.json"
		}
	}

	opt := option.WithCredentialsFile(credentialsPath)
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	if projectID == "" {
		projectID = "stay-watch-a616f"
	}

	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("Cannot initialize firebase app: %v\n", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Cannot initialize firebase auth: %v\n", err)
	}

	firebaseApp = app
	firebaseAuth = authClient
}
