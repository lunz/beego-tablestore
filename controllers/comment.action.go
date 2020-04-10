package controllers

import (
	m "beego-tablestore/models"
	"beego-tablestore/otslib"
	"encoding/json"
	"fmt"
)

/***************************************
*	Post
{
	"ProdId":"123",
	"Content":"I feel greate",
}
"CommentId":<int>, if value is provided, this is an update action
****************************************/
func (c *CommentController) AddOrUpdateContent() {

	var comment m.Comment
	json.Unmarshal(c.Ctx.Input.RequestBody, &comment)
	//TODO: validation

	out, _ := json.Marshal(comment)
	fmt.Println(string(out))

	if commentId, err := otslib.InsertOrUpdate(comment); err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = commentId
	}
	c.ServeJSON()
}

/***************************************
*	Get
****************************************/
func (c *CommentController) Get() {

	prodId, commentId := c.ReadParams()

	resp, err := otslib.GetComment(prodId, commentId)
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

/***************************************
*	Delete
****************************************/
func (c *CommentController) Delete() {

	prodId, commentId := c.ReadParams()

	if err := otslib.DeleteComment(prodId, commentId); err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = "Deleted"
	}
	c.ServeJSON()
}

/***************************************
*	Batch Get page by page
****************************************/
func (c *CommentController) BatchGetByProdId() {

	prodId, commentId, pageSize := c.ReadPageParams()

	if batchResp, commentId, err := otslib.GetMultiCommentsByPage(prodId, commentId, pageSize); err != nil {
		c.Data["json"] = err
	} else {
		ret := make(map[string]interface{})
		ret["Rows"] = batchResp
		ret["Next"] = commentId
		c.Data["json"] = ret
	}
	c.ServeJSON()
}
