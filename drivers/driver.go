// [4] Frameworks & Drivers
// Controllerを呼び出し、echoの設定を行う

package drivers

import (
	"clean-architecture-learning/adapters/controllers"
	"context"

	"github.com/labstack/echo/v4"
)

type User interface {
	ServeUsers(ctx context.Context, address string)
}

type UserDriver struct {
	echo       *echo.Echo
	controller controllers.User
}

func NewUserDriver(echo *echo.Echo, controller controllers.User) User {
	return &UserDriver{
		echo:       echo,
		controller: controller,
	}
}

func (driver *UserDriver) ServeUsers(ctx context.Context, address string) {
	driver.echo.POST("/users", driver.controller.AddUser(ctx))
	driver.echo.GET("/users", driver.controller.GetUser(ctx))
	driver.echo.Logger.Fatal(driver.echo.Start(address))
}
