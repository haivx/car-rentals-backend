package repo

import (
	"car-rentals-backend/model"
	"car-rentals-backend/util"
	"errors"
	"time"
)

func Login(req *model.Login) (post *model.Login, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("tiêu đề không được để trống")
	}

	if err != nil {
		return nil, err
	}

	return req, nil
}

func Register(req *model.Register) (user *model.User, err error) {
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
