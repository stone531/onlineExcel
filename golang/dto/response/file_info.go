package response

import "time"

type MetaFileInfo struct {
	FileId      string 			`json:"file_Id"` //文档唯一标识

	FileName 	string 			`json:"name"` //文件名
	Author 		string 			`json:"author"` //文档的作者

	Version 	int 			`json:"version"`

	CreateStime time.Time 		`json:"create_stime"`  //时间戳
	Modifier    string    		`json:"modifier"`
	UpdateStime time.Time 		`json:"update_stime"`
}