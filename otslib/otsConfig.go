package otslib

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/astaxie/beego"
)

var (
	AccessKeyId      = beego.AppConfig.String("ots::AccessKeyId")
	AccessKeySecret  = beego.AppConfig.String("ots::AccessKeySecret")
	InternetEndPoint = beego.AppConfig.String("ots::InternetEndPoint")
	InstanceName     = beego.AppConfig.String("ots::InstanceName")
	CommentTableName = "DummyTest"
)

func CreateClient() *tablestore.TableStoreClient {
	return tablestore.NewClient(InternetEndPoint, InstanceName, AccessKeyId, AccessKeySecret)
}

func UnmarshalComment(resp *tablestore.Row, v interface{}) error {

	var sb strings.Builder
	sb.WriteString(`{`)

	for _, col := range resp.PrimaryKey.PrimaryKeys {

		if reflect.TypeOf(col.Value).Name() == "int64" {
			sb.WriteString(fmt.Sprintf(`"%s":%d,`, col.ColumnName, col.Value))
		} else {
			sb.WriteString(fmt.Sprintf(`"%s":"%s",`, col.ColumnName, col.Value))
		}
	}

	for _, col := range resp.Columns {
		if reflect.TypeOf(col.Value).Name() == "int64" {
			sb.WriteString(fmt.Sprintf(`"%s":%d,`, col.ColumnName, col.Value))
		} else {
			sb.WriteString(fmt.Sprintf(`"%s":"%s",`, col.ColumnName, col.Value))
		}
	}

	jstr := strings.TrimRight(sb.String(), ",") + "}"

	return json.Unmarshal([]byte(jstr), v)
}
