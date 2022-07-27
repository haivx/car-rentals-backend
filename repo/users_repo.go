package repo

import (
	"car-rentals-backend/model"
	"car-rentals-backend/util"
	"errors"
	"time"
)

func GetAllUser() (users []model.User, err error) {
	err = DB.Model(&users).Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id string) (user *model.User, err error) {
	user = &model.User{
		Id: id,
	}

	err = DB.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(req *model.CreateUser) (user *model.User, err error) {
	if req.FullName == "" || req.Phone == "" || req.Email == "" {
		return nil, errors.New("must not empty")
	}

	user = &model.User{
		Id:        util.NewID(),
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
	}
	_, err = DB.Model(user).WherePK().Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	token := "3432"
	user.Token = token
	return user, nil
}
