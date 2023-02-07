package excel

import (
	"ark-online-excel/common"
	"ark-online-excel/dao"
	"ark-online-excel/dto"
	"ark-online-excel/dto/request"
	"ark-online-excel/dto/response"
	"ark-online-excel/global"
	"ark-online-excel/logger"
	"ark-online-excel/models"
	"ark-online-excel/utils"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

/**
* @program: src
* @description: v1版本的报表生成
* @author: lirl
* @create 2023-1-14 14:26
**/

// 最新文件的生成
var sheetname = "Sheet1"


// 生成文件的主逻辑
func GenerateSheetFile(req *request.CreateMateData) (*response.CreateMateData, error) {

	serverTestFile(req.Name, req.Author)

	fileId := common.GetNewDocId()
	mateData := &models.ExcelMeta {
		Author:		req.Author,
		FileId:     fileId,
		FileName:	req.Name,

		Cell:     common.InitExcelCellData,
		Data:     common.InitExcelData,
		RowData:  common.InitExcelRowData,
		Modifier: req.Author,
	}

	//写入数据如果慢，需要goroutine
	_,err := dao.AddOnlineExcelFileData(mateData)
	if err != nil {
		logger.ZapLogger.Error("创建文件失败 err:", zap.Error(err))
		return &response.CreateMateData{},err
	}

	res := &response.CreateMateData{
		FileId:fileId,
	}

	return res ,nil
}

func serverTestFile(name, author string) error {
	/**
	创建文件
	*/
	fileHandle,err := utils.Init_file(sheetname)
	if err != nil {
		logger.ZapLogger.Error("serverTestFile err:", zap.Error(err))
		return err
	}

	/**
	设置行列宽度
	*/
	var cellData dto.SheetCells
	json.Unmarshal([]byte(common.InitExcelCellData),&cellData)
	fmt.Println("GenerateSheetFile cell:%s",cellData)
	if err := SetColsAndRowslength(fileHandle,cellData,sheetname);err!=nil{
		logger.ZapLogger.Error("serverTestFile err:", zap.Error(err))
		return err
	}

	/**
	循环写入样式元数据和单元格值
	*/
	var excelData []dto.SheetDataGroup
	json.Unmarshal([]byte(common.InitExcelData),&excelData)
	fmt.Println("GenerateSheetFile data:%s",excelData)
	if err := SetBlockStyleAndValueEx(fileHandle,excelData,sheetname,author);err!=nil {
		logger.ZapLogger.Error("serverTestFile err:", zap.Error(err))
		return err
	}

	file_dir := global.GetBaseFilePath() +"/"+ name //+ constants.Name_time_mark + strconv.Itoa(int(time.Now().Unix())) + ".xlsx"
	// 文件写入磁盘
	//file_dir := req.Name + constants.Name_time_mark +  strconv.Itoa(int(time.Now().Unix())) + ".xlsx"
	if err := fileHandle.SaveAs(file_dir);err!=nil {
		logger.ZapLogger.Error("serverTestFile err:", zap.Error(err))
		return err
	}

	return nil
}


func QuerySheetFile(req *request.QueryMetaFile) (*response.QueryMetaFiles, error){

	//metas,err := dao.QueryMetaByConditions(req.FileName,req.User, req.Version)
	metas,err := dao.GetAllMetaFile(1000)
	if err != nil {
		logger.ZapLogger.Error("QuerySheetFile err:", zap.Error(err))
		return nil, err
	}

	var res []response.SheetFile

	for _,meta := range metas {
		tmpMetaFile := response.SheetFile {
			FileId: meta.FileId,
			Version: meta.Version,
			FileName:meta.FileName,
			CreateTime: meta.CreateStime.String(),
		}

		res = append(res, tmpMetaFile)
	}

	return &response.QueryMetaFiles {
		FileList: res,
	}, nil
}

func QuerySheetMetaData(fileId string) (*response.OpenMetaFile, error){

	meta,err := dao.GetMetaDataById(fileId)
	if err != nil {
		logger.ZapLogger.Error("QuerySheetMetaData err:", zap.Error(err))
		return nil, err
	}

	return &response.OpenMetaFile {
		Time : meta.CreateStime.String(),
		Author:meta.Author,
		Name:meta.FileName,
		Cell:meta.Cell,
		Data:meta.Data,
		RawData:meta.RowData,
	}, nil
}

func DeleteSheetFile(fileId string) error {
	return  dao.DelFileById(fileId)
}

func getFileNewVersion(fileId string) (int,error) {
	return dao.GetMetaFileVersion(fileId)
}

func SaveSheetMetaData(req *request.SaveMateData) error {

	jCellData,err := json.Marshal(req.Cell)
	if err != nil {
		logger.ZapLogger.Error("SaveSheetMetaData err:", zap.Error(err))
		return err
	}

	jData,err := json.Marshal(req.Data)
	if err != nil {
		logger.ZapLogger.Error("SaveSheetMetaData err:", zap.Error(err))
		return err
	}

	_,err = dao.UpdateFileInfo(req.FileId, req.User,string(jCellData),string(jData),req.RowData)
	if err != nil {
		logger.ZapLogger.Error("SaveSheetMetaData err:", zap.Error(err))
		return err
	}

	return nil
}

func GetFileInfo(fileId string) (*response.MetaFileInfo,error) {

	metadata,err := dao.GetMetaDataById(fileId)
	if err != nil {
		logger.ZapLogger.Error("GetFileInfo err:", zap.Error(err))
		return nil, err
	}

	return &response.MetaFileInfo{
		FileId:metadata.FileId,
		FileName:metadata.FileName,
		Author:metadata.Author,
		CreateStime:metadata.CreateStime,
		UpdateStime:metadata.UpdateStime,
		Modifier:metadata.Modifier,
		Version:metadata.Version,
	},nil

}