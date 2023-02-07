package request

type CreateMateData struct {
	Name 	string `json:"name" gorm:"name"` //文件名
	Author 	string `json:"author" gorm:"author"` //文档的作者
}


