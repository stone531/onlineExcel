package controllers

import (
	"fmt"
	"time"

	webs "ark-online-excel/service/ws"
)

type ExcelController struct {
	BaseController
}

var g_count int
func (c *ExcelController) JoinExcel() {

	user := c.GetString("user")
	fmt.Println("ws JoinExcel params_bk :",user)

	g_count++
	username := fmt.Sprintf("li-test-%d",g_count)

	//webs.JoinShareExcelRoom(username,c.Ctx.ResponseWriter, c.Ctx.Request)
	c.ShareExcel(username,"test.xls")
	//c.SetClientError("集群获取管理员时发生错误")
}


func (c *ExcelController) ShareExcel(userName,fileName string) {

	if userName == "" || fileName == "" {
		c.SetClientError("ShareExcel params_bk err")
		return
	}

	shareObj := webs.FindFileObj(fileName)
	if shareObj == nil {
		newObj :=webs.NewShareRoom()

		//add map
		webs.AddMap(fileName, newObj)

		c.SetData("open file success ")
		return
	}

	//只有文档共享人数大于2人及以上，需要加入共享队列
	if shareObj.AddShareCount() == 1 {
		//shareObj.NewShareRoom()
		go shareObj.StartShareExcel(fileName)
		time.Sleep(500)
	}

	shareObj.JoinShareExcelRoom(userName, fileName, c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	shareObj.AddShareCount()

}

