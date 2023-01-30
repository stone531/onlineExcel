package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(ClusterAnalysis))
}

type ClusterAnalysis struct {
	Id          int64         `json:"id"`
	ClusterId   int64         `json:"cluster_id"`
	SummaryId   int64         `json:"summary_id"`
	DetectType  int           `json:"detect_type"`
	MetricName  string        `json:"metric_name"`
	MetricValue string        `json:"metric_value"`
	Summary     string        `json:"summary"`
	Suggestion  string        `json:"suggestion"`
	Creator     string        `json:"creator"`
	CreateStime time.Time     `orm:"auto_now_add;type(datetime)" json:"create_stime"`
	Modifier    string        `json:"modifier"`
	UpdateStime time.Time     `orm:"auto_now_add;type(datetime)" json:"update_stime"`
	AlarmsList  []interface{} `orm:"-" json:"alarms_list"`
	Content     string        `orm:"-" json:"content"`
	MetricCount int           `orm:"-" json:"metric_count"`
	Final       bool          `orm:"-" json:"final"`
}

func (a *ClusterAnalysis) TableName() string {
	return "cluster_analysis"
}
