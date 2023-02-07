package ws

import (
	"ark-online-excel/common"
	"ark-online-excel/dto/response"
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"ark-online-excel/logger"
	"container/list"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"time"
)

const (
	// 允许等待的写入时间
	readWait = 1 * time.Minute

	// Time allowed to read the next pong message from the peer.
	pongWait = 1 * time.Minute

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

)

//public struct
type WebSocket struct {
	beego.Controller
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}


type ExcelShareObj struct {
	subscribe 	chan 	Subscriber
	unsubscribe chan 	string
	publish 	chan  	common.Message
	subscribers *list.List
	shareCount	int  //共享次数
	runFlag     bool //当前房间是否可用
	//shareUser   string//共享人
}

func NewShareRoom () *ExcelShareObj {
	return &ExcelShareObj{
		subscribe:	make(chan Subscriber, common.WebSocket_Max_User_Count),
		unsubscribe:make(chan string, common.WebSocket_Max_User_Count),
		publish:	make(chan common.Message, common.WebSocket_Content_Max_Size),
		//quickCh:    make(chan int,10),
		subscribers:list.New(),
		shareCount: 0,
		runFlag:	true,
	}
}

func (es *ExcelShareObj)AddShareCount () int {
	if es == nil {
		return 0
	}
	es.shareCount ++

	return es.shareCount
}

func (es *ExcelShareObj)SendMessage(types int, user string, msg interface{}, fileName string) common.Message {
	return common.Message{types, user, time.Now().Unix(), msg, fileName}
}

func (es *ExcelShareObj)StartShareExcel(fileName string) {
	logger.ZapLogger.Info("enter shareExcel ....")

	fmt.Println("StartShareExcel enter....",es.subscribe)
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		fmt.Println("StartShareExcel room end....",es.subscribe)
		ticker.Stop()
		es.shareCount = 0
		//webs.DelMap(fileName)
		es.notifyAllUserClose()
		es.runFlag = false
	}()

	for {
		select {
		case sub := <-es.subscribe:
			//if !es.isUserExist(es.subscribers, sub.Name) {
			es.subscribers.PushBack(sub)
			es.publish <- es.SendMessage(common.Msg_Op_Type_User_Online, sub.Name, fmt.Sprintf("user %s Join success,ws连接成功",sub.Name),fileName)
			//}

		case message := <-es.publish:
			es.broadcastWebSocket(message)

		case unsub := <-es.unsubscribe:
			for sub := es.subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					es.subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
					}
					es.publish <- es.SendMessage(common.Msg_Op_Type_User_Offline, unsub, fmt.Sprintf("user %s Join success,ws关闭连接",unsub),fileName)
					break
				}
			}

			//es.quickCh <- es.subscribers.Len()

		case <-ticker.C:
			// 出现超时情况
			fmt.Println("StartShareExcel check time out ==")
			//无需服务端发送ping，需要客户端每几分钟发送pong到服务端即可
			//es.broadcastHeartbeat()

			if es.subscribers.Len() <= 1 {
				fmt.Printf("StartShareExcel filename:%s no user room exit,count:%d\n",fileName,es.subscribers.Len())
				logger.ZapLogger.Info("disconnect ",zap.String("StartShareExcel filename:%s no user server exit",fileName))
				time.Sleep(time.Second*2)
				return
			}
		}

	}
}


func (es *ExcelShareObj)broadcastWebSocket(msg common.Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		logger.ZapLogger.Error("broadcastWebSocket Fail to marshal event:", zap.Error(err))
		return
	}

	for sub := es.subscribers.Front(); sub != nil; sub = sub.Next() {

		//如果消息是自己发出的无需给自己发送
		if sub.Value.(Subscriber).Name == msg.User {
			continue
		}

		ws := sub.Value.(Subscriber).Conn
		if ws != nil {

			//todo 有用户退出，是否需要通知其它用户，退出的消息
			err = ws.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				es.unsubscribe <- sub.Value.(Subscriber).Name //通知关闭该用户
				logger.ZapLogger.Error("broadcastWebSocket user send msg exception err:", zap.Error(err))
			}
		} else{
			logger.ZapLogger.Error("broadcastWebSocket get connect handler err:", zap.String("user ",sub.Value.(Subscriber).Name))
		}
	}
}

func (es *ExcelShareObj)broadcastHeartbeat() {

	for sub := es.subscribers.Front(); sub != nil; sub = sub.Next() {

		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				logger.ZapLogger.Error("broadcastHeartbeat user send Ping msg exception err:", zap.Error(err))
				//关闭当前连接
				es.disConnect(ws)
				return
			}
			fmt.Println("broadcastHeartbeat user send Ping success")

		} else{
			logger.ZapLogger.Error("broadcastWebSocket get connect handler err:", zap.String("user ",sub.Value.(Subscriber).Name))
		}
	}
}

//关闭所有用户连接，房间退出
func (es *ExcelShareObj)notifyAllUserClose() {
	for sub := es.subscribers.Front(); sub != nil; sub = sub.Next() {

		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			//ws.WriteMessage(websocket.CloseMessage, []byte{})

			es.disConnect(ws)

			fmt.Println("notifyAllUserClose close success")
		}
	}
}

func (es *ExcelShareObj)isUserExist(subscribers *list.List, user string) bool {

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			fmt.Println("isUserExist true", user)
			return true
		}
	}

	fmt.Println("isUserExist false", user)
	return false
}


//有新用户加入共享excel 房间 用户数至少有两个以上才创建房间。
//http连接转成tcp长连接
func (es *ExcelShareObj)JoinShareExcelRoom(username,fileName string,writer http.ResponseWriter, req *http.Request, initMsg *response.OpenMetaFile) {
	ws, err := websocket.Upgrade(writer, req, nil, common.WebSocket_send_Max_Len, common.WebSocket_send_Max_Len)
	if val, ok := err.(websocket.HandshakeError); ok {
		logger.ZapLogger.Error("JoinShareExcelRoom join Error", zap.String(" convert tcp Upgrade ",val.Error()))
		return
	} else if err != nil {
		logger.ZapLogger.Error("JoinShareExcelRoom join err Error", zap.Error(err))
		return
	}

	if initMsg != nil {
		es.sendInitMetaData(ws,initMsg)
	}

	//订阅者加入
	es.subscribe <- Subscriber{Name: username, Conn: ws}
	logger.ZapLogger.Info("user ",zap.String("join success",username))

	//发生异常时，通知关闭该用户
	defer func() {
		es.unsubscribe <- username
		fmt.Println("JoinShareExcelRoom  ====00 user,end read",username)
		logger.ZapLogger.Info("JoinShareExcelRoom exit", zap.String("join success",username))
	}()

	es.setConnSetting(ws,username)

	for {
		//开启监听已连接的用户
		_, metaData, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("JoinShareExcelRoom user disconnect .",username)
			logger.ZapLogger.Info("user ",zap.String("user ",username),zap.Error(err))
			return
		}
		fmt.Println("JoinShareExcelRoom read msg:",string(metaData),username)

		//对接收的消息检查是否有效
		var umsg common.Message
		err = json.Unmarshal([]byte(metaData), &umsg)
		if err != nil {
			fmt.Println("JoinShareExcelRoom read msg Unmarshal err:",err,username)
			logger.ZapLogger.Info("user ",zap.String("user ",username),zap.Error(err),zap.Any("err msg msg :",umsg))
			continue
		}

		if umsg.Type < common.Msg_Op_Type_Min || umsg.Type > common.Msg_Op_Type_Max {
			fmt.Println("JoinShareExcelRoom read msg type err:",umsg.Type, username)
			logger.ZapLogger.Info("err msg discard :",zap.String("user:",username),zap.Int("err type:",umsg.Type))
			continue
		}

		if umsg.Type == common.Msg_Op_Type_Pong {
			es.vueHeartBeat(ws)
			continue
		}

		es.PrintAllUser()
		fmt.Println("JoinShareExcelRoom start send msg",string(metaData), username)
		es.publish <- es.SendMessage(common.Msg_Op_Type_Content, username, umsg.Content,fileName)
	}

}

func (es *ExcelShareObj)PrintAllUser() {
	for sub := es.subscribers.Front(); sub != nil; sub = sub.Next() {
		fmt.Println("current user :", sub.Value.(Subscriber).Name)
	}
}
func (es *ExcelShareObj)sendInitMetaData(conn *websocket.Conn,metaData *response.OpenMetaFile) {

	msg := common.Message{common.Msg_Op_Type_File_Init, metaData.Author, time.Now().Unix(), metaData, metaData.Name}

	es.publish <- msg
}

func (es *ExcelShareObj)setConnSetting(conn *websocket.Conn,user string) {
	//设置接收消息超时时间
	conn.SetReadDeadline(time.Now().Add(readWait))

	//conn.SetWriteDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		fmt.Println("接收心跳响应<<",user)
		conn.SetReadDeadline(time.Now().Add(readWait))
		return nil
	})
}

func (es *ExcelShareObj)disConnect(conn *websocket.Conn) {
	conn.Close()
}

func (es *ExcelShareObj)vueHeartBeat(conn *websocket.Conn) {
	fmt.Println("接收vue心跳响应<<")
	conn.SetReadDeadline(time.Now().Add(readWait))
}