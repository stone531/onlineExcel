package services

import (
	"ark-online-excel/models"
)

const (
	EnvLevelProd = "prod"
	EnvLevelTest = "test"
	EnvLevelAll  = "all"
)

const (
	SystemCreateName = "admin"
)

var LocalIp string

type BaseService struct {
	LoginUser *models.LoginUser
}
