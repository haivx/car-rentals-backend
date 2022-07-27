package controller

import (
	"car-rentals-backend/model"
	"car-rentals-backend/repo"

	"github.com/kataras/iris/v12"
)

func GetAllNews(ctx iris.Context) {
	news, err := repo.GetAllNews()
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
	}

	ctx.JSON(news)
}

func GetNewsById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	news, err := repo.GetNewsById(id)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(news)
}

func CreateNews(ctx iris.Context) {
	req := new(model.News)

	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := repo.CreateNews(req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(createdUser)
}
