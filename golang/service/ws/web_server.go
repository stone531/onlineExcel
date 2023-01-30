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
	quickCh     chan    int// subscribe 的长度
	subscribers *list.List
	shareCount	int  //共享次数
	//shareUser   string//共享人
}

func NewShareRoom () *ExcelShareObj {
	return &ExcelShareObj{
		subscribe:	make(chan Subscriber, 10),
		unsubscribe:make(chan string, 10),
		publish:	make(chan common.Message, 10),
		quickCh:    make(chan int,10),
		subscribers:list.New(),
		shareCount: 0,
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

	defer func() {
		fmt.Println("StartShareExcel end....",es.subscribe)
	}()
	for {
		select {
		case sub := <-es.subscribe:
			if !es.isUserExist(es.subscribers, sub.Name) {
				es.subscribers.PushBack(sub)
				es.publish <- es.SendMessage(common.Msg_Op_Type_User_Online, sub.Name, fmt.Sprintf("user %s Join success,ws连接成功",sub.Name),fileName)
			}

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

			es.quickCh <- es.subscribers.Len()

		case quick := <- es.quickCh :
			//es.subscribers.Len() == 0:
			//文档在线人数小于1时退出共享文档
			fmt.Printf("StartShareExcel111 filename:%s no user server exit,count:%d\n",fileName,quick)
			if quick <= 1 {
				fmt.Printf("StartShareExcel filename:%s no user server exit,count:%d\n",fileName,quick)
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
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				es.unsubscribe <- sub.Value.(Subscriber).Name //通知关闭该用户
			}
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
	ws, err := websocket.Upgrade(writer, req, nil, 1024, 1024)
	if val, ok := err.(websocket.HandshakeError); ok {
		fmt.Println("JoinShareExcelRoom 00")
		logger.ZapLogger.Error("JoinShareExcelRoom join Error", zap.String(" convert tcp Upgrade ",val.Error()))
		return
	} else if err != nil {
		fmt.Println("JoinShareExcelRoom 01")
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
		logger.ZapLogger.Info("JoinShareExcelRoom exit", zap.String("join success",username))
	}()

	for {
		//开启监听已连接的用户
		fmt.Println("JoinShareExcelRoom wait input msg...")
		_, metaData, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("JoinShareExcelRoom user disconnect .",username)
			logger.ZapLogger.Info("user ",zap.String("user ",username),zap.Error(err))
			return
		}
		fmt.Println("JoinShareExcelRoom read msg2",string(metaData))
		es.publish <- es.SendMessage(common.Msg_Op_Type_Content, username, string(metaData),fileName)

	}
}

func (es *ExcelShareObj)sendInitMetaData(conn *websocket.Conn,metaData *response.OpenMetaFile) {

	msg := common.Message{common.Msg_Op_Type_File_Init, metaData.Author, time.Now().Unix(), metaData, metaData.Name}

	es.publish <- msg
	//if conn.WriteMessage(websocket.TextMessage, msg) != nil
	//return conn.WriteJSON(msg)
}