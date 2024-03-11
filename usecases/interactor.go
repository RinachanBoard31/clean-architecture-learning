// [2] Application Business Rules
// InputPortの実態を実装
// 技術的要素を省いた部分でinterface(./portに記載している定義)を呼び出すことで、usecaseの外側にある技術的要素を隠蔽している

package interactors

import (
	"context"
	"firestore_clean/entities"
	"firestore_clean/usecases/ports"
)

// ./portに定義しているinterface
type UserInteractor struct {
	OutputPort ports.UserOutputPort
	Repository ports.UserRepository
}

func NewUserInputPort(outputPort ports.UserOutputPort, repository ports.UserRepository) ports.UserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *UserInteractor) AddUser(ctx context.Context, user *entities.User) error {
	users, err := u.Repository.AddUser(ctx, user)
	if err != nil {
		return &u.OutputPort.outputError(err)
	}

	return &u.OutputPort.OutputUser(users)
}

func (u *UserInteractor) GetUser(ctx context.Context) error {
	users, err := u.Repository.GetUser(ctx)
	if err != nil {
		return &u.OutputPort.OutputError(err)
	}

	return &u.OutputPort.OutputUser(users)
}
