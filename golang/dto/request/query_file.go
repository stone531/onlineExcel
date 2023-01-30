package request

type QueryMetaFile struct {
	FileName 		string     `json:"file_name"`
	Version 		string	   `json:"version"`
	User 			string     `json:"user"`
	RemoteDir 		string 	   `json:"remote_dir"`//扩展字段后期可能需要打开不同目录下相同文件
}
