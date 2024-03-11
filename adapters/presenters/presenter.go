// [3] Interface Adapter
// OutputPortの実態を定義する
// 技術的要素(echo, firestore)の実装を含む

package presenters

import (
	"clean-architecture-learning/entities"
	"clean-architecture-learning/usecases/ports"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserPresenter struct {
	ctx echo.Context
}

func NewUserOutputPort(ctx echo.Context) ports.UserOutputPort {
	return &UserPresenter{
		ctx: ctx,
	}
}

// 3. 全てのユーザを返す
func (u *UserPresenter) OutputUser(users []*entities.User) error {
	return u.ctx.JSON(http.StatusOK, users)
}

// 3. エラーを返す
func (u *UserPresenter) OutputError(err error) error {
	log.Fatal(err)
	return u.ctx.JSON(http.StatusInternalServerError, err)
}
