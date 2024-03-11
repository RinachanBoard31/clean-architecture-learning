// [2] Application Business Rules
// 1. echoでユーザデータJSONを受け取る -> ユーザデータを受け取る
// 2. firestoreにユーザデータを追加する -> 何かしらにユーザデータを追加する
// 3. echoで全てのユーザデータJSON/エラーを返す -> 全てのユーザデータJSON/エラーを返す
// usecaseでは、技術的な要素を除いたビジネスルールを記述する

package ports

import (
	"context"
	"firestore_clean/entities"
)

// 1. ユーザデータ受け取り
type UserInputPort interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context) error
}

// 2. 何かしらにユーザデータを追加する
type UserRepository interface {
	AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error)
	GetUser(ctx context.Context) ([]*entities.User, error)
}

// 3. 全てのユーザデータJSON/エラーを返す
type UserOutputPort interface {
	OutputUser([]*entities.user) error
	outputError(error) error
}
