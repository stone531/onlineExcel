package response

type OpenMetaFile struct {
	Time 	string 		`json:"time" gorm:"time"` //时间戳
	Name 	string 		`json:"name" gorm:"name"` //文件名
	Author 	string 		`json:"author" gorm:"author"` //文档的作者
	Cell    string      `json:"cell"` //单元格扩展样式
	Data    string 		`json:"data"` //单元格元数据
	RawData string      `json:"rawdata"`
}