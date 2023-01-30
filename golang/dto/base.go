package dto

// 文件的公用数据
type PublicSheetParams struct {
	Time string `json:"time" gorm:"time"` //时间戳
	Name string `json:"name" gorm:"name"` //文件名
	Author string `json:"author" gorm:"author"` //文档的作者
}

type ApiInfo struct {
	// api的信息
	Type string `json:"type" gorm:"type"`
	Host string `json:"host" gorm:"host"`
	Port string `json:"port" gorm:"port"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"'`
	DBName string `json:"dbname" gorm:"dbname"`
	TableName string `json:"tablename" gorm:"tablename"`
}

//单元格扩展的元数据
type SheetCells struct {
	Cols []SheetCols `json:"cols"` //扩展的单元格列
	Rows []SheetRows `json:"rows"` //扩展的单元格行
}

type SheetCols struct {
	Index string `json:"index"`
	Width string `json:"width"`
}

type SheetRows struct {
	Index string `json:"index"`
	Height string `json:"height"`
}

//单元格元数据
type SheetDataGroup struct {
	Merge []string `json:"mergeOrNot"` //合并单元格
	Style Styles `json:"style"` //样式属性
	Text string `json:"text"` //文本值
}

//样式属性
type Styles struct {
	Color string `json:"color"`
	BgColor string `json:"bgcolor"`
	Align string `json:"align"`
	Valign string `json:"valign"`
	Font SheetFont `json:"font"`
	Underline bool `json:"underline"`
	Border SheetBorder `json:"border"` //边框样式
	TextWrap bool `json:"textwrap"` //是否自动换行
}

type SheetFont struct {
	Name string `json:"name"`
	Size float64 `json:"size"` //字体大小
	Bold bool `json:"bold"` //是否加粗
	Italic bool `json:"italic"` //是否斜体字
}

type SheetBorder struct {
	Bottom []string `json:"bottom"`
	Top []string `json:"top"`
	Left []string `json:"left"`
	Right []string `json:"right"`
	//"bottom": ["thick", "#fe0000"],
	//"top": ["thick", "#fe0000"],
	//"left": ["thick", "#fe0000"],
	//"right": ["thick", "#fe0000"]
}


// 表名信息
type SheetTableReq struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName string `json:"dbname"`
	TableName string `json:"tablename"`
}

// 表元数据的ID
type SheetRawdataReq struct {
	ID int64 `json:"id"`
}

// 获取文件操作历史请求
type SheetHistoryReq struct {
	//TableName string `json:"table_name"`
	Offset int32 `json:"offset"`
}

// 样式的工具类
type StyleJson struct {
	Name   string
	Params []byte
}