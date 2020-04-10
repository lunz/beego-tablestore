package otslib

import (
	m "beego-tablestore/models"
	"fmt"
	"strconv"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func BuildPutChangeMeta(c m.Comment) *tablestore.PutRowChange {

	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = CommentTableName
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	if c.CommentId == 0 {
		putRowChange.SetReturnPk()
	}

	meta := new(tablestore.PrimaryKey)
	meta.AddPrimaryKeyColumn("ProdId", c.ProdId)
	if c.CommentId > 0 {
		meta.AddPrimaryKeyColumn("CommentId", c.CommentId)
	} else {
		meta.AddPrimaryKeyColumnWithAutoIncrement("CommentId")
	}
	putRowChange.PrimaryKey = meta

	putRowChange.AddColumn("Content", c.Content)

	return putRowChange
}

func InsertOrUpdate(commentStorage m.Comment) (int64, error) {

	request := new(tablestore.PutRowRequest)
	request.PutRowChange = BuildPutChangeMeta(commentStorage)

	client := CreateClient()
	resp, err := client.PutRow(request)

	if resp.PrimaryKey.PrimaryKeys != nil {
		col := resp.PrimaryKey.PrimaryKeys[1]
		commentId, _ := strconv.ParseInt(fmt.Sprintf(`%d`, col.Value), 10, 64)
		return commentId, err
	}

	return 0, err
}
