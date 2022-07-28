package controller

import (
	"car-rentals-backend/model"
	"car-rentals-backend/repo"

	"golang.org/x/crypto/bcrypt"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	req := new(model.Login)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}
	currentUser, err := repo.GetUserByUserName(req.Username)

	if currentUser == nil || err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, "user's not existed")
		return
	}
	notValidPassword := bcrypt.CompareHashAndPassword([]byte(string(currentUser.Password)), []byte(req.Password))

	if notValidPassword != nil {
		ResponseErr(ctx, iris.StatusBadRequest, "username's not existed or wrong password")
		return
	}

	loginInfo, err := repo.Login(currentUser)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(loginInfo)
}
