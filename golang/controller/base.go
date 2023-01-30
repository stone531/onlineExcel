package controllers

import (
	"ark-online-excel/logger"
	"ark-online-excel/models"
	"ark-online-excel/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go.uber.org/zap"
	"strings"
)

const (
	MaxRequestBodyLength = 1024 * 64
	TokenStr             = "tokenId"
	ErrNoRowsString      = "请求出错，请确认资源存在"
)

type BaseController struct {
	beego.Controller
	result    models.JsonRet
	loginUser *models.LoginUser
	//RecordLog  bool
	//AuthType   string
	UserString string
	//UserInfo   map[string]interface{} //{ "account": "ss", "employeeId": "12222","username": "张三"}
	//IsAdmin    bool
	//TeamIds    []int64 // 所属团队
	//Teams      []models.TeamRole
}

var routerWhiteList = map[string]bool{
	"/api/v1/health":               true,
	"/api/v1/cluster/getYamlByKey": true,
	//"/api/v1/ws/diagnosis":         true,
}

var wildcardRouterList = []string{"/api/v1/log-manage/app/docker/paths/"}

func wildcardRouterJudger(path string) bool {
	for _, v := range wildcardRouterList {
		if strings.Contains(path, v) {
			return true
		}
	}
	return false
}

//用户登陆信息获取并验证
func (bc *BaseController) Prepare() {


}

func (bc *BaseController) Finish() {
	bc.Data["json"] = bc.result
	bc.ServeJSON()
}

func (bc *BaseController) SetErrorWithData(err string, data interface{}) {
	if strings.Contains(err, orm.ErrNoRows.Error()) {
		err = ErrNoRowsString
	}

	bc.result.Err = err
	bc.result.Status = false
	bc.result.Data = data
	bc.result.Code = 500
}

func (bc *BaseController) SetError(err string) {
	if strings.Contains(err, orm.ErrNoRows.Error()) {
		err = ErrNoRowsString
	}

	bc.result.Err = err
	bc.result.Status = false
	bc.result.Code = 500
}

func (bc *BaseController) SetClientError(err string) {
	if strings.Contains(err, orm.ErrNoRows.Error()) {
		err = ErrNoRowsString
	}

	bc.result.Err = err
	bc.result.Status = false
	bc.result.Code = 400
}

func (bc *BaseController) SetAuthorizationError(err string) {
	if strings.Contains(err, orm.ErrNoRows.Error()) {
		err = ErrNoRowsString
	}

	bc.result.Err = err
	bc.result.Status = false
	bc.result.Code = 401
}

func (bc *BaseController) SetData(data interface{}) {
	bc.result.Data = data
	bc.result.Status = true
	bc.result.Err = ""
	bc.result.Code = 200
}

func (bc *BaseController) CheckErr(err error) bool {
	if err != nil {
		logger.ZapLogger.Error("CheckErr", zap.Error(err))
		return false
	} else {
		return true
	}
}

func (bc *BaseController) GetPageInfo() (page int, size int, withcount bool, err bool) {
	page, _ = bc.GetInt("page", 0)
	size, _ = bc.GetInt("size", 10)
	if page < 1 || size < 1 {
		bc.SetError("page 或者 pagesize 传参错误")
		err = true
	}
	withcount, _ = bc.GetBool("withcount", false)
	return
}

func (bc *BaseController) ParseBody(req interface{}) (interface{}, error) {
	body := bc.Ctx.Input.RequestBody
	length := len(body)
	if length > utils.MaxBodyLen {
		return nil, utils.ErrMaxPostBody
	}

	err := json.Unmarshal(body, req)
	if err != nil {
		err = fmt.Errorf("post body解析出错: %s", err.Error())
	}
	return req, err
}
