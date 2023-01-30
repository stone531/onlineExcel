package response

import "ark-online-excel/dto"

type MateData struct {
	dto.PublicSheetParams
	//Api ApiInfo `json:"api" gorm:"api"` // 文件下载路径
	Cell dto.SheetCells `json:"cell"` //单元格扩展样式
	Data []dto.SheetDataGroup `json:"data"` //单元格元数据
	RawData string `json:"rawdata"`
}