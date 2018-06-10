package firebase

import (
	"log"

	"github.com/fidelicash/fc-api/util"
	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"google.golang.org/api/option"
)

func Conn() (*db.Client, error) {
	dURL := util.GetOSEnvironment("FDATABASE_URL", "https://fidelicash-c6245.firebaseio.com")
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: dURL,
	}

	// Fetch the service account key JSON file contents
	// opt := option.WithCredentialsFile("serviceAccountKey.json")
	opt := option.WithCredentialsFile("")
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
