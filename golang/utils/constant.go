package utils

import (
	"errors"
)

const (
	EventStatusAlert   = 1 //告警中
	EventStatusRecover = 2 //已恢复
)

const (
	AlertLevelDisaster = 0
	AlertLevelSerious  = 2
	AlertLevelWarning  = 4
	AlertLevelNotice   = 6

	AlertLevelDisasterStr = "紧急"
	AlertLevelSeriousStr  = "重要"
	AlertLevelWarningStr  = "次要"
	AlertLevelNoticeStr   = "提示"
)

const (
	LogTypeDeny                = "屏蔽"
	LogTypeKnow                = "知悉"
	LogTypeRecover             = "恢复"
	LogTypeSendUpgrade         = "发送告警升级"
	LogTypeSendUpgradeRecovery = "发送告警升级恢复"
	LogTypeSendAlert           = "发送告警"
	LogTypeSendAlertRepeat     = "发送重复告警"
	LogTypeSendAlertCall       = "发送回调"
)

const Recovery = "resolved"

const (
	AlertAnnotationsModule = "alert_module"
	AlertAnnotationsRule   = "alert_rule"
	RelationTagKey         = "__relation_tag_key"

	AlertStatus = "auto_monitor_alert_status"
)

const DbMonitor = "monitor"

const MaxBodyLen = 1024 * 64

//MatchType
//0: =
//1: !=
//2:正则
const (
	MatchTypeEqual    = 0
	MatchTypeNotEqual = 1
	MatchTypeRegex    = 2
)

//type RepeatSendAlert struct {
//	key           string
//	event         *entity.AlertEvent
//	strategyOnce  *models.StrategyOnce
//	alertStrategy *models.AlertStrategy
//}

const (
	alertRepeatlastSendTimeKey       = "autoinsight_alertrepeat_"
	alertRepeatlastSendTimeKeyExpire = 3600 * 24 * 2

	judgeAddr = "http://monitor-insight-judge%s.openapi.corpautohome.com/"
)

const (
	RuleValue   = "value"
	RuleTitle   = "title"
	RuleContent = "content"
	ExprCheck   = "%sapi/v1/query?time=%d&query="

	CloudMonitorVersionKey = "cloud_monitor_version_key"

	//RulesVersionKey      = "autoinsight_rules_version"
	VersionKeySimpleRules  = "cloud_monitor_simple_rules_version"
	VersionKeyComplexRules = "cloud_monitor_complex_rules_version"

	RulesField           = "rules"
	DataSourceVersionKey = "autoinsight_datasource_version"
	DataSourceField      = "datasource"

	MetricReportGroupVersionKey = "insight_metric_report_group_version"
	MetricReportGroupField      = "insight_metric_report_group"
)

const (
	SEND_TYPE_DING        = "钉钉"
	SEND_TYPE_SMS         = "短信"
	SEND_TYPE_AUTO        = "汽车人"
	SEND_TYPE_CALL        = "电话"
	SEND_TYPE_DINGROBOTS  = "钉钉机器人"
	SEND_TYPE_DINGSMS     = "钉钉|短信"
	SEND_TYPE_DINGSMSAUTO = "钉钉|短信|汽车人"
	SEND_TYPE_SMSAUTO     = "短信|汽车人"
	SEND_TYPE_DINGAUTO    = "钉钉|汽车人"
	DefaultReceiveType    = SEND_TYPE_DING
)

const CALLNUMLIMIT = 10

// 告警类别
const (
	CONTAINERBASE = iota + 1 // 1.容器基础
	CONTAINERHTTP = iota + 1 // 2.容器HTTP状态码请求
	CONTAINERJVM  = iota + 1 // 3.容器JVM
	HOSTBASE      = iota + 1 // 4.主机基础
	HOSTJVM       = iota + 1 // 5.主机JVM
	URL           = iota + 1 // 6.URL告警
	LOGMONITOR    = iota + 1 // 7.日志监控
	CUSTOM        = iota + 1 // 8.自定义监控
)

const (
	AlertRepeatlastSendTimeKey       = "autoinsight_alertrepeat_"
	AlertRepeatlastSendTimeKeyExpire = 3600 * 24 * 2

	JudgeAddr = "http://monitor-insight-judge%s.openapi.corpautohome.com/"
)

var (
	ErrNotPermission = errors.New("没有权限操作改记录")
	ErrNotExist      = errors.New("数据已经不存在")
	ErrMaxPostBody   = errors.New("post body exceed the maximum value")
)

// 返回常量字符串
const (
	OkStr   = "Ok"
	NullStr = ""
)

const (
	AllEnableTeam = 530
)

// 消息中心的发送通道代码
const (
	Ding        = 100
	Sms         = 200
	Email       = 300
	AutoOa      = 400
	AutoAOne    = 410
	AutoMonitor = 420
)

const Sender = "system"

// 发送告警情况
const (
	SendFail = iota // 0.发送失败
	SendOk   = iota // 1.发送成功
)

const (
	SendFailStr = "失败"
	SendOkStr   = "成功"
)

const (
	ModuleSourceUrl        = "Url"
	ModuleSourceGatewayUrl = "GatewayUrl"
	ModuleSourceAutoLog    = "AutoLog"
	ModuleSourceCustom     = "Custom"
)

const (
	RunModeTest = "test"
	RunModeDev  = "dev"
	RunModeProd = "prod"
)

const (
	AlertObjectHost = "host"
	AlertObjectUrl  = "url"
)

const (
	//容器磁盘规则ID
	PodDiskRuleDefault85Id  = 18
	PodDiskRuleDefault90Id  = 19
	PodDiskRuleStateful85Id = 20
	PodDiskRuleStateful90Id = 21
)

const (
	DataSourceCustom = "prometheus-custom"
	DataSourceTest   = "prometheus-test"
)

const (
	MonitorWriteQPS          = "监控指标写入QPS"
	MonitorReadRequest       = "查询请求总量"
	MonitorAppCount          = "接入应用总数"
	MonitorCustomMetricCount = "自定义监控指标数"
)
