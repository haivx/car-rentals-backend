package repo

import (
	"car-rentals-backend/model"
	"car-rentals-backend/util"
	"errors"
	"time"
)

func GetAllUser() (users []model.User, err error) {
	if users == nil {
		return nil, nil
	}
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

func GetUserByUserName(username string) (user *model.User, err error) {
	user = &model.User{
		Username: username,
	}
	err = DB.Model(user).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(req *model.CreateUser) (user *model.User, err error) {
	if req.Username == "" || req.Phone == "" || req.Email == "" {
		return nil, errors.New("must not empty")
	}

	user = &model.User{
		Id:        util.NewID(),
		Username:  req.Username,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
		Password:  req.Password,
	}
	_, err = DB.Model(user).WherePK().Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	return user, nil
}
