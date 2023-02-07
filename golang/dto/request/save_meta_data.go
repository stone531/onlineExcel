package request

import "ark-online-excel/dto"

type SaveMateData struct {
	FileId  string 				`json:"file_id" gorm:"file_id"` //文件Id
	User 	string 				`json:"user" gorm:"user"` //用户

	Cell    dto.SheetCells       `json:"cell"` //单元格扩展样式
	Data    []dto.SheetDataGroup `json:"data"` //单元格元数据
	RowData string               `json:"row_data"`
}
