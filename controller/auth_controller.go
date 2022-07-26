package controller

import (
	"car-rentals-backend/model"
	"car-rentals-backend/repo"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	req := new(model.Login)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	loginInfo, err := repo.Login(req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(loginInfo)
}
