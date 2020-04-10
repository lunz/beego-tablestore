package routers

import (
	"beego-tablestore/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/table/create", &controllers.CommentController{}, "post:CreateTable"),
		beego.NSRouter("", &controllers.CommentController{}, "post:AddOrUpdateContent"),
		beego.NSRouter("/:prodId/:commentId", &controllers.CommentController{}, "get:Get;delete:Delete"),
		beego.NSRouter("/batch/:prodId", &controllers.CommentController{}, "get:BatchGetByProdId"),
		beego.NSRouter("/batch/:prodId/:commentId", &controllers.CommentController{}, "get:BatchGetByProdId"),
	)
	beego.AddNamespace(ns)
}
