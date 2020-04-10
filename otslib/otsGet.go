package otslib

import (
	m "beego-tablestore/models"
	"fmt"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func GetComment(prodId string, commentId int64) (m.Comment, error) {

	meta := new(tablestore.PrimaryKey)
	meta.AddPrimaryKeyColumn("ProdId", prodId)
	meta.AddPrimaryKeyColumn("CommentId", commentId)

	criteria := new(tablestore.SingleRowQueryCriteria)
	criteria.PrimaryKey = meta

	getRowRequest := new(tablestore.GetRowRequest)
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = CommentTableName
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1

	client := CreateClient()
	getResp, err := client.GetRow(getRowRequest)

	row := tablestore.Row{PrimaryKey: &getResp.PrimaryKey, Columns: getResp.Columns}

	var comment m.Comment
	if err := UnmarshalComment(&row, &comment); err != nil {
		return comment, err
	}

	if err != nil {
		fmt.Println("GetComment failed with error:", err)
	} else {
		fmt.Println("Comment column col0 result is ", getResp.Columns[0].ColumnName, getResp.Columns[0].Value)
	}
	return comment, err
}

func GetMultiCommentsByPage(prodId string, commentId int64, pageSize int32) ([]m.Comment, int64, error) {

	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	rangeRowQueryCriteria.TableName = CommentTableName
	rangeRowQueryCriteria.MaxVersion = 1
	rangeRowQueryCriteria.Limit = pageSize
	rangeRowQueryCriteria.Direction = tablestore.FORWARD

	if commentId > 0 { //pagination
		pageStartPK := new(tablestore.PrimaryKey)
		pageStartPK.AddPrimaryKeyColumn("ProdId", prodId)
		pageStartPK.AddPrimaryKeyColumn("CommentId", commentId)

		rangeRowQueryCriteria.StartPrimaryKey = pageStartPK

	} else {
		startPK := new(tablestore.PrimaryKey)
		startPK.AddPrimaryKeyColumn("ProdId", prodId)
		startPK.AddPrimaryKeyColumnWithMinValue("CommentId")
		rangeRowQueryCriteria.StartPrimaryKey = startPK
	}

	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumn("ProdId", prodId)
	endPK.AddPrimaryKeyColumnWithMaxValue("CommentId")
	rangeRowQueryCriteria.EndPrimaryKey = endPK

	getRangeRequest := &tablestore.GetRangeRequest{}
	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria

	client := CreateClient()
	getRangeResp, err := client.GetRange(getRangeRequest)
	if err == nil && len(getRangeResp.Rows) > 0 {

		nextStartCommentId := int64(0)
		rowCount := len(getRangeResp.Rows)
		comments := make([]m.Comment, rowCount)

		for i := 0; i < rowCount; i++ {
			row := getRangeResp.Rows[0]
			var comment m.Comment
			if err := UnmarshalComment(row, &comment); err != nil {
				fmt.Println("-- failed to unmarshall comment - ", err)
			}
			comments[i] = comment
		}
		if getRangeResp.NextStartPrimaryKey != nil {
			for _, col := range getRangeResp.NextStartPrimaryKey.PrimaryKeys {
				if col.ColumnName == "CommentId" {
					if val, ok := col.Value.(int64); ok {
						nextStartCommentId = val
					}
				}
			}
		}
		return comments, nextStartCommentId, nil
	}

	return nil, 0, err
}
