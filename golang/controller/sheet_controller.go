package controllers

import (
	"fmt"
	"time"

	"ark-online-excel/dto/request"
	//"ark-online-excel/models/params_bk"
	"ark-online-excel/service/excel"

	webs "ark-online-excel/service/ws"
)

type SheetController struct {
	BaseController
}

// 生成报表
func (c *SheetController)CreateSheetData() {

	req := &request.MateData{}
	_, err := c.ParseBody(req)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}

	fileId, err := excel.GenerateSheetFile(req)
	if err!=nil {
		c.SetClientError(err.Error())
		return
	}

	c.SetData(fileId)
}

func (c *SheetController)QuerySheetFile() {
	req := &request.QueryMetaFile{}
	_, err := c.ParseBody(req)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}

	if req.FileName == "" || req.User == "" {
		c.SetClientError(fmt.Sprintf("SheetController QuerySheetFile params_bk empty"))
		return
	}

	metaFileListObj, err := excel.QuerySheetFile(req)
	if err!=nil {
		c.SetClientError(err.Error())
		return
	}

	c.SetData(metaFileListObj)
}


func (c *SheetController)OpenSheetMateData() {

	fileId :=c.GetString("fileId")
	if fileId == "" {
		c.SetClientError("OpenSheetMateData get params_bk err file_id empty")
		return
	}

	user := c.GetString("user")
	if user == "" {
		c.SetClientError("OpenSheetMateData get params_bk err user empty")
		return
	}

	fmt.Println("OpenSheetMateData 01:",fileId,user)

	c.shareExcel(fileId,user)
}


func (c *SheetController) shareExcel(fileId,user string) {

	fmt.Println("shareExcel 00")
	shareObj := webs.FindFileObj(fileId)
	if shareObj == nil {
		newObj :=webs.NewShareRoom()

		//add map
		webs.AddMap(fileId, newObj)

		mFile,err := excel.QuerySheetMetaData(fileId)
		if err != nil {
			c.SetClientError(err.Error())
			return
		}

		c.SetData(mFile)
		return
	}

	//只有文档共享人数大于2人及以上，需要加入共享队列
	if shareObj.AddShareCount() == 1 {
		//shareObj.NewShareRoom()
		go shareObj.StartShareExcel(fileId)
		time.Sleep(500)
	}

	shareObj.AddShareCount()

	//给当前用户发送文档初始数据
	mFile,err := excel.QuerySheetMetaData(fileId)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}
	c.SetData(mFile)

	//开启长连接
	shareObj.JoinShareExcelRoom(user, fileId, c.Ctx.ResponseWriter, c.Ctx.Request, mFile)

}

// 根据ID获取表的模版数据
func (c *SheetController)GetTableRawData() {
	//var req params_bk.SheetRawdataReq
	//if err := c.BindJSON(&req);err!=nil {
	//	log.Fatal(err)
	//	response.ResponseError(c,400,"参数绑定错误",err)
	//	return
	//}
	//data,err := services.GetExcelRawDatas(req.ID)
	//if err!= nil {
	//	log.Fatal(err)
	//	response.ResponseError(c,500,err.Error(),err)
	//	return
	//}
	//response.ResponseSuccess(c,200,"获取成功",data)
}


// 获取数据源的字段信息
func (c *SheetController)GetSheetTableMeta() {
	//var req params_bk.SheetTableReq
	//if err := c.BindJSON(&req);err!=nil {
	//	log.Fatal(err)
	//	response.ResponseError(c,400,"参数绑定错误",err)
	//	return
	//}
	//data,err := services.GetTableMetaInfo(req.TableName)
	//if err!= nil {
	//	log.Fatal(err)
	//	response.ResponseError(c,500,"生成失败",err)
	//	return
	//}
	//response.ResponseSuccess(c,200,"生成成功",data)
}

// 获取表格的操作历史
func (c *SheetController)GetSheetHistory() {
	//var req params_bk.SheetHistoryReq
	//if err := c.BindJSON(&req);err!=nil {
	//	log.Fatal(err)
	//	response.ResponseError(c,400,"参数绑定错误",err)
	//	return
	//}
	//data,err := services.GetSheetHistory(req)
	//if err!= nil {
	//	log.Fatal(err)
	//	response.ResponseError(c,500,"获取失败",err)
	//	return
	//}
	//response.ResponseSuccess(c,200,"获取成功",data)
}

func (c *SheetController)SaveSheetData() {

}

func (c *SheetController)DelSheetFile() {
	fileId :=c.GetString("fileId")
	if fileId == "" {
		c.SetClientError("OpenSheetMateData get params_bk err file_id empty")
		return
	}

	err := excel.DeleteSheetFile(fileId)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}

	c.SetData("success")
}