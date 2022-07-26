package controller

import (
	"car-rentals-backend/model"
	"car-rentals-backend/repo"

	"github.com/kataras/iris/v12"
)

func GetAllUser(ctx iris.Context) {
	users, err := repo.GetAllUser()
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
	}

	ctx.JSON(users)
}

func GetUserById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	user, err := repo.GetUserById(id)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(user)
}

func CreateUser(ctx iris.Context) {
	req := new(model.CreateUser)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := repo.CreateUser(req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(createdUser)
}
