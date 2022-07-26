package repo

import (
	"car-rentals-backend/model"
	"errors"
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
