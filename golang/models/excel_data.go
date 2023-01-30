package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(ExcelMeta))
}

type ExcelMeta struct {
	Id 			int64 			`json:"id"`
	FileId      string 			`json:"file_Id"` //文档唯一标识

	FileName 	string 			`json:"name"` //文件名
	Author 		string 			`json:"author"` //文档的作者

	//Api 		[]byte 			`json:"api" gorm:"api"` // 文件下载路径
	Cell    	string 			`json:"cell" orm:"cell"`
	Data    	string 			`json:"data" orm:"data"`
	RowData 	string 			`json:"rowdata"`
	Version 	int 			`json:"version"`

	CreateStime time.Time 		`json:"create_stime" orm:"auto_now_add;type(datetime)"`  //时间戳
	Modifier    string    		`json:"modifier"`
	UpdateStime time.Time 		`json:"update_stime" orm:"auto_now_add;type(datetime)"`
}

func (a *ExcelMeta) TableName() string {
	return "oe_meta"
}
