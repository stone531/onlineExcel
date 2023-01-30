package excel

import (
	"ark-online-excel/common"
	"ark-online-excel/dto/response"
	"encoding/json"
	"log"

	"ark-online-excel/dao"
	"ark-online-excel/dto/request"
	"ark-online-excel/global"
	//"ark-online-excel/middlewares_bk/constants"
	"ark-online-excel/models"
	//model "ark-online-excel/models/db"
	"ark-online-excel/utils"
)

/**
* @program: src
* @description: v1版本的报表生成
* @author: lirl
* @create 2023-1-14 14:26
**/

// 最新文件的生成
var sheetname = "Sheet1"
// sqlx连接器
//var opsqlx model.OpSqlxExcelMetaDao
// gorm连接器
//var opgorm model.OpGormExcelMetaDao

// 生成文件的主逻辑
func GenerateSheetFile(req *request.MateData) (string, error) {
	/**
	创建文件
	 */
	f,err := utils.Init_file(sheetname)
	if err != nil {
		log.Fatal(err)
		return "error",err
	}

	/**
	设置行列宽度
	*/
	if err := SetColsAndRowslength(f,req.Cell,sheetname);err!=nil{
		log.Fatal(err)
		return "error",err
	}

	/**
	循环写入样式元数据和单元格值
	 */
	if err := SetBlockStyleAndValueEx(f,req.Data,sheetname,req.Author);err!=nil {
		log.Fatal(err)
		return "error",err
	}

	// 协程写入数据库，这里要添加异常处理
	//var OpExcelDao model.OpGormExcelMetaDao
	//go OpExcelDao.WriteData(global.DBOrmEngine,req)

	jCellData,_:=json.Marshal(req.Cell)
	jData,_:=json.Marshal(req.Data)

	file_id := common.GetNewDocId()
	mateData := &models.ExcelMeta {
		Author:		req.Author,
		FileId:     file_id,
		FileName:	req.Name,

		Cell:     string(jCellData),
		Data:     string(jData),
		RowData:  req.RawData,
		Modifier: req.Author,
	}

	//写入数据如果慢，需要goroutine
	_,err = dao.AddOnlineExcelFileData(mateData)
	if err != nil {
		log.Fatal(err)
		return "db error",err
	}

	file_dir := global.GetBaseFilePath() +"/"+ req.Name //+ constants.Name_time_mark + strconv.Itoa(int(time.Now().Unix())) + ".xlsx"
	// 文件写入磁盘
	//file_dir := req.Name + constants.Name_time_mark +  strconv.Itoa(int(time.Now().Unix())) + ".xlsx"
	if err := f.SaveAs(file_dir);err!=nil {
		log.Fatal(err)
		return "error",err
	}

	//url := global.GetFileUrl() + "/" + file_dir
	return file_id ,nil
}


func QuerySheetFile(req *request.QueryMetaFile) (*response.QueryMetaFiles, error){

	metas,err := dao.QueryMetaByConditions(req.FileName,req.User, req.Version)
	if err != nil {
		return nil, err
	}

	var res []response.SheetFile

	for _,meta := range metas {
		tmpMetaFile := response.SheetFile {
			FileId: meta.FileId,
			Version: meta.Version,
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