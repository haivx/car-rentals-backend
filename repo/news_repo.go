package repo

import (
	"car-rentals-backend/model"
	"car-rentals-backend/util"
)

func GetAllNews() (news []model.News, err error) {

	err = DB.Model(&news).Select()

	if err != nil {
		return nil, err
	}

	return news, nil
}

func GetNewsById(id string) (news *model.News, err error) {
	news = &model.News{
		Id: id,
	}
	err = DB.Model(news).WherePK().Select()

	if err != nil {
		return nil, err
	}

	return news, nil
}

func CreateNews(req *model.News) (news *model.News, err error) {
	news = &model.News{
		Id:      util.NewID(),
		Title:   req.Title,
		Summary: req.Summary,
		Image:   req.Image,
		Author:  req.Author,
		Content: req.Content,
	}
	_, err = DB.Model(news).WherePK().Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	return news, nil
}
