package router

import (
	"car-rentals-backend/controller"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {
	v1 := app.Party("/api/v1")
	{
		v1.Use(iris.Compression)

		v1.Post("/auth/login", controller.Login)

		v1.Post("/users", controller.CreateUser)
		v1.Get("/users", controller.GetAllUser)
		v1.Get("/users/:id", controller.GetUserById)
		v1.Post("/users/register", controller.Register)

		v1.Post("/news", controller.CreateNews)
		v1.Get("/news", controller.GetAllNews)
		v1.Get("/news/:id", controller.GetNewsById)
	}

}
