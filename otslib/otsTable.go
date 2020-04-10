package otslib

import (
	"fmt"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func CreateTable() error {

	fmt.Println("Start creating table - ", CommentTableName)

	createTableRequest := new(tablestore.CreateTableRequest)

	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 1
	createTableRequest.TableOption = tableOption

	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createTableRequest.ReservedThroughput = reservedThroughput

	meta := new(tablestore.TableMeta)
	meta.TableName = CommentTableName
	meta.AddPrimaryKeyColumn("ProdId", tablestore.PrimaryKeyType_STRING)                                      // prodid as partition key
	meta.AddPrimaryKeyColumnOption("CommentId", tablestore.PrimaryKeyType_INTEGER, tablestore.AUTO_INCREMENT) //row key for each comment
	meta.AddDefinedColumn("Like", tablestore.DefinedColumn_INTEGER)                                           //prod LIKE  count to replace author id
	meta.AddDefinedColumn("ReplyTo", tablestore.DefinedColumn_INTEGER)                                        //if this is top comment
	meta.AddDefinedColumn("FromEmail", tablestore.DefinedColumn_STRING)                                       //comment user id
	meta.AddDefinedColumn("Content", tablestore.DefinedColumn_STRING)                                         //comments text
	meta.AddDefinedColumn("Imgs", tablestore.DefinedColumn_STRING)                                            //user uploaded images

	createTableRequest.TableMeta = meta

	//create
	client := CreateClient()
	_, err := client.CreateTable(createTableRequest)

	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}

	return err
}

func IsTableExisted(tableName string) (bool, error) {

	client := CreateClient()
	tables, err := client.ListTable()
	if err != nil {
		fmt.Println("Failed to list tables")
		return false, err
	} else {
		fmt.Println("available tables are ")
		for _, table := range tables.TableNames {
			fmt.Println("table name: ", table)
			if table == tableName {
				return true, nil
			}
		}
		return false, nil //the default table doesn't exist
	}
}
