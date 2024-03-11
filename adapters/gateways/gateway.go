// [3] Interface Adapter
// DBの実装(firestore)

package gateways

import (
	"context"
	"encoding/json"
	"errors"
	"firestore_clean/entities"
	"firestore_clean/usecases/ports"
	"fmt"

	"cloud.google.com/go/firestore"
)

type FirestoreClientFactory interface {
	NewClient(ctx context.Context) (*firestore.Client, error)
}

type UserGateway struct {
	clientFactory FirestoreClientFactory
}

// usecase/portに依存している
func NewUserRepository(clientFactory FirestoreClientFactory) ports.UserRepository {
	return &UserGateway{
		clientFactory: clientFactory,
	}
}

// eintityに依存している
func (gateway *UserGateway) AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	client, err := gateway.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed AddUser NewClient: %v", err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{}{
		"empNumber":  user.EmpNumber,
		"department": user.Department,
	})

	if err != nil {
		return nil, fmt.Errorf("failed AddUser Set: %v", err)
	}

	return getUsers(ctx, client)
}

func (gateway *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {
	client, err := gateway.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers NewClient: %v", err)
	}
	defer client.Close()

	return getUsers(ctx, client)
}

func getUsers(ctx context.Context, client *firestore.Client) ([]*entities.User, error) {
	allData := client.Collection("users").Documents(ctx)

	docs, err := allData.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers GetAll: %v", err)
	}

	users := make([]*entities.User, 0)
	for _, doc := range docs {
		u := new(entities.User)
		err = mapToStruct(doc.Data(), &u)
		if err != nil {
			return nil, fmt.Errorf("failed GetUsers mapToStruct: %v", err)
		}
		u.Name = doc.Ref.ID
		users = append(users, u)
	}

	return users, nil
}

// map -> 構造体の変換
func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
