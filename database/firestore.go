// [4] Frameworks & Drivers
// gatewayで宣言しているFirestoreClientFactoryを実装した

package database

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type MyFirestoreClientFactory struct{}

func (f *MyFirestoreClientFactory) NewClient(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("credentialsFile.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
