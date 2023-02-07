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

	req := &request.CreateMateData{}
	_, err := c.ParseBody(req)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}

	if req.Author == "" || req.Name == "" {
		c.SetClientError(fmt.Sprintf("CreateSheetData params empty, auth=%s,name=%s",req.Author,req.Name))
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

	if req.User == "" {
		c.SetClientError(fmt.Sprintf("SheetController QuerySheetFile author empty"))
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

	//给当前用户发送文档初始数据
	mFile,err := excel.QuerySheetMetaData(fileId)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}
	//c.SetData(mFile)

	//开启长连接
	shareObj.JoinShareExcelRoom(user, fileId, c.Ctx.ResponseWriter, c.Ctx.Request, mFile)

}

func (c *SheetController)SaveSheetData() {
	req := &request.SaveMateData{}
	_, err := c.ParseBody(req)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}

	if req.User == "" || req.FileId == "" {
		c.SetClientError(fmt.Sprintf("SaveSheetData params empty, user=%s,fileId=%s",req.User,req.FileId))
		return
	}

	if len(req.Data) == 0 {
		c.SetClientError("SaveSheetData data empty,no data need save")
		return
	}

	err = excel.SaveSheetMetaData(req)
	if err!=nil {
		c.SetClientError(err.Error())
		return
	}

	c.SetData(fmt.Sprintf("save fileId:%s success",req.FileId))
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

func (c *SheetController)GetFileInfo() {
	fileId :=c.GetString("fileId")
	if fileId == "" {
		c.SetClientError("OpenSheetMateData get params_bk err file_id empty")
		return
	}

	res,err := excel.GetFileInfo(fileId)
	if err != nil {
		c.SetClientError(err.Error())
		return
	}

	c.SetData(res)
}