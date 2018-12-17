package controller

import (
	"github.com/kataras/iris"
	"../models/db"
	"../models"
	)

type UserController struct {
	BaseController
}

func (userController *UserController) Add(ctx iris.Context){
	user := models.User{Name:"wendao", Age: 30, Sex:"M"}
	db.Db().Create(&user)
	ctx.JSON([]string{"message"})
	ctx.Next()
}

