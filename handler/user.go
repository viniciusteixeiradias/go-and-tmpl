package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/vdias/got/model"
	"github.com/vdias/got/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "vinicius@gmail.com",
	}

	return render(c, user.Show(u))
}
