package common


const (
	Msg_Op_Type_Add 			= 1 //用户增加数据
	Msg_Op_Type_Del 			= 2 //用户删除数据
	Msg_Op_Type_Update 			= 3 //用户更新数据
)

//type 1:用户上线，2.文档内容数据  3. 用户下线. 4.文档初始数据,5. 检查用户心跳ack
const (
	Msg_Op_Type_Min				= 0
	Msg_Op_Type_User_Online 	= 1
	Msg_Op_Type_Content 		= 2
	Msg_Op_Type_User_Offline 	= 3
	Msg_Op_Type_File_Init 		= 4
	Msg_Op_Type_Ping		 	= 5
	Msg_Op_Type_Pong		 	= 6


	Msg_Op_Type_Max				= Msg_Op_Type_Pong
)


type Message struct {
	Type      int 			`json:"type"`
	User      string 		`json:"user"`
	Timestamp int64 		`json:"timestamp"`
	Content   interface{} 	`json:"content"`
	FileName  string    	`json:"fileName"`
}



