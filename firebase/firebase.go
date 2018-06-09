package firebase

import (
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"google.golang.org/api/option"
)

func Conn() (*db.Client, error) {

	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://fidelicash-c6245.firebaseio.com",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return client, nil
}

func Get(client *db.Client, reference string) (interface{}, error) {
	// As an admin, the app has access to read and write all data, regradless of Security Rules
	// ref := client.NewRef("restricted_access/secret_document")
	ctx := context.Background()
	ref := client.NewRef(reference)
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	return data, nil
}
