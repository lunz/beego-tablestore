package otslib

import (
	"fmt"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func DeleteComment(prodId string, commentId int64) error {

	meta := new(tablestore.PrimaryKey)
	meta.AddPrimaryKeyColumn("ProdId", prodId)
	meta.AddPrimaryKeyColumn("CommentId", commentId)

	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = CommentTableName
	deleteRowReq.DeleteRowChange.PrimaryKey = meta
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)

	// clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	// deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)

	client := CreateClient()
	_, err := client.DeleteRow(deleteRowReq)

	if err != nil {
		fmt.Println("getrow failed with error:", err)
	} else {
		fmt.Println("delete row finished")
	}

	return err
}
