package common


// 文件的公用数据
type PublicSheetParams struct {
	Time 	string `json:"time" gorm:"time"` //时间戳
	Name 	string `json:"name" gorm:"name"` //文件名
	Author 	string `json:"author" gorm:"author"` //文档的作者
}

