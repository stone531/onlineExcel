package response

type QueryMetaFiles struct {
	FileList []SheetFile `json:"file_list"`
}


type SheetFile struct {
	FileId 		string      `json:"file_id"`
	FileName 	string      `json:"file_name"`
	Version 	int	   		`json:"version"`
	CreateTime	string 		`json:"create_time"`
}