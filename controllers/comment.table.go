package controllers

import (
	"beego-tablestore/otslib"
	"fmt"
)

func (c *CommentController) CreateTable() {

	ok, err := otslib.IsTableExisted(otslib.CommentTableName)

	if !ok {
		if err != nil {
			fmt.Println("err  when checking if table exists")
			c.Data["json"] = err //todo: replace with 401
		} else {
			if err := otslib.CreateTable(); err == nil {
				fmt.Println("err == null when creating table")
				c.Data["json"] = "Table is created."
			} else {
				fmt.Println("err != null when creating table")
				c.Data["json"] = err
			}
		}
	} else {
		c.Data["json"] = "Default table is already existed."
	}

	c.ServeJSON()
}
