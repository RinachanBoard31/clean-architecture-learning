// [3] Interface Adapter
// InputPort, OutputPort, Repositoryを組み立てて、InputPortを実行する

package controllers

import (
	"clean-architecture-learning/adapters/gateways"
	"clean-architecture-learning/entities"
	"clean-architecture-learning/usecases/ports"
	"context"

	"github.com/labstack/echo/v4"
)

type User interface {
	AddUser(ctx context.Context) func(c echo.Context) error
	GetUser(ctx context.Context) func(c echo.Context) error
}

type OutputFactory func(echo.Context) ports.UserOutputPort
type InputFactory func(ports.UserOutputPort, ports.UserRepository) ports.UserInputPort
type RepositoryFactory func(gateways.FirestoreClientFactory) ports.UserRepository

type UserController struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
	clientFactory     gateways.FirestoreClientFactory
}

func NewUserController(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory, clientFactory gateways.FirestoreClientFactory) User {
	return &UserController{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
		clientFactory:     clientFactory,
	}
}

func (u *UserController) AddUser(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(entities.User)
		err := c.Bind(user)

		if err != nil {
			return err
		}

		return u.NewInputPort(c).AddUser(ctx, user)
	}
}
func (u *UserController) GetUser(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		return u.NewInputPort(c).GetUser(ctx)
	}
}

func (u *UserController) NewInputPort(c echo.Context) ports.UserInputPort {
	outputPort := u.outputFactory(c)
	repository := u.repositoryFactory(u.clientFactory)
	return u.inputFactory(outputPort, repository)
}
