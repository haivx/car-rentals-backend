package repo

import (
	"car-rentals-backend/model"
	"car-rentals-backend/services"
	"car-rentals-backend/util"
	"errors"
	"time"
)

func Login(req *model.User) (user *model.User, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("tiêu đề không được để trống")
	}

	token, err := services.GenerateToken(req)

	if err != nil {
		return nil, err
	}

	req.Token = token

	return req, err
}

func Register(req *model.CreateUser) (user *model.User, err error) {
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, errors.New("data must not empty")
	}

	user = &model.User{
		Id:        util.NewID(),
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}
	_, err = DB.Model(user).WherePK().Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	return user, nil
}
