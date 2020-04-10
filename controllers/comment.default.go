package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (base *CommentController) RespondError(statusCode int, message ...string) {

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.Output.SetStatus(statusCode)
	base.Data["json"] = message
	base.ServeJSON()
	base.StopRun()
}

func (base *CommentController) ReadParams() (string, int64) {

	prodId := base.ReadProdId()

	commentId, err := strconv.ParseInt(base.Ctx.Input.Param(":commentid"), 10, 64)
	if err != nil || commentId == 0 {
		base.RespondError(400, "Comment id is not found")
	}

	return prodId, commentId
}

func (c *CommentController) ReadPageParams() (string, int64, int32) {

	prodId := c.ReadProdId()

	commentId, err := strconv.ParseInt(c.Ctx.Input.Param(":commentid"), 10, 64)
	if err != nil {
		commentId = 0
	}

	pageSize, err := c.GetInt32("top")
	if err != nil {
		pageSize = 1
	}
	return prodId, commentId, pageSize
}

func (base *CommentController) ReadProdId() string {

	prodId := base.Ctx.Input.Param(":prodid")
	if prodId == "" {
		base.Ctx.ResponseWriter.Write([]byte(""))
		base.RespondError(400, "Prod id is not found")
	}
	return prodId
}
