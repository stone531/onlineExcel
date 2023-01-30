package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(ApiLog))
}

type ApiLog struct {
	Id          int       `json:"id"`
	Method      string    `json:"method"`
	Url         string    `json:"url"`
	Address     string    `json:"address"`
	Body        string    `orm:"type(text)" json:"body"`
	Creator     string    `json:"creator"`
	CreateStime time.Time `orm:"auto_now_add;type(datetime)" json:"create_stime"`
	Modifier    string    `json:"modifier"`
	UpdateStime time.Time `orm:"auto_now_add;type(datetime)" json:"update_stime"`
}

func (a *ApiLog) TableName() string {
	return "api_log"
}
