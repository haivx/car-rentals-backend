package repo

import (
	"car-rentals-backend/model"
	"car-rentals-backend/util"
	"errors"
	"fmt"
	"os"
	"time"
)

func GetAllUser() (users interface{}, err error) {
	fmt.Println(os.Getenv("POSTGRES_USER"))
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"))
	fmt.Println(os.Getenv("POSTGRES_DB"))
	// err = DB.Model(&users).Select()

	if err != nil {
		return nil, err
	}

	return os.Getenv("POSTGRES_USER"), nil
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

func CreateUser(req *model.CreateUser) (post *model.CreateUser, err error) {
	if req.FullName == "" || req.Phone == "" || req.Email == "" {
		return nil, errors.New("Must not empty")
	}

	user := &model.User{
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

	return req, nil
}
