package request

type ExcelMetaData struct {
	Time string `json:"time" gorm:"time"` //时间戳
	Name string `json:"name" gorm:"name"` //文件名
	Author string `json:"author" gorm:"author"` //文档的作者
	data []byte //json 只传差异数据，服务器只转发数据
}
