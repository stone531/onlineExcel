package common


// 文件的公用数据
type PublicSheetParams struct {
	Time 	string `json:"time" gorm:"time"` //时间戳
	Name 	string `json:"name" gorm:"name"` //文件名
	Author 	string `json:"author" gorm:"author"` //文档的作者
}

var InitExcelCellData  string = `{"cols":[{"index":"A","width":"20"},{"index":"C","width":"20"}],"rows":[{"index":"A2","width":"52"}]}`

var InitExcelData string = `[{"mergeOrNot":["A1","A2"],"style":{"color":"#777777","BgColor":"#777777","align":"cccc","font":{"name":"16px","size":26.0,"bold":false,"italic":true},"underline":true,"border":{"bottom":["thick","#fe0000"],"top":["thick","#fe0000"],"left":["thick","#fe0000"],"right":["thick","#fe0000"]},"textwrap":false},"text":""}]`

var InitExcelRowData string = ``

const (

	//websocket 传输最大值
	WebSocket_send_Max_Len = 309760

	//允许最大用户数
	WebSocket_Max_User_Count = 50

	//允许最大消息积压条数
	WebSocket_Content_Max_Size = 1024
)


