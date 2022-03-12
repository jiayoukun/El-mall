package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"test15/pkg/e"
)

//发送信息的类型
type SendMsg struct {
	Code int `json:"code"`
	Type int `json:"type"`
	Content string `json:"content"`
}
//回复的信息
type ReplyMsg struct {
	From string `json:"from"`
	Code int `json:"code"`
	Content string `json:"content"`
}
type Broadcast struct {
	Client 	*User
	Message []byte
	Type 	int
}
type ClientManager struct {
	Clients 	map[string]*User
	Broadcast 	chan *Broadcast
	Reply 		chan *User
	Register 	chan *User
	Unregister  chan *User
}
var Manager  = ClientManager{
	Clients:    make(map[string]*User), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *User),
	Reply:      make(chan *User),
	Unregister: make(chan *User),

}
//用户类
type User struct {
	ID	string
	SendID string
	Socket *websocket.Conn
	Send chan []byte
	UserLock sync.RWMutex
}
//用户管理

//消息转json
type Message struct {
	Sender string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content string `json:"content,omitempty"`
}
//type Server struct {
//	User *User
//	SendID string
//	Broadcast *Broadcast
//	OnlineMap map[string]*User
//	Message chan string
//}
var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Start(){
	for{
		log.Println("<---监听管道通信--->")
		select {
		case conn := <-Manager.Register: // 建立连接
			log.Printf("建立新连接: %v", conn.ID)
			Manager.Clients[conn.ID] = conn
			replyMsg := &ReplyMsg{
				Code: 200,
				Content: "用户ID:"+conn.ID+"已连接至服务器",
			}
			log.Printf("replyMsg",replyMsg)
			msg, _ := json.Marshal(replyMsg)

			//conn.Send <-"已连接至服务器"
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister: // 断开连接
			log.Printf("连接失败:%v", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    300,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
			//广播信息
		case broadcast := <-Manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false // 默认对方不在线

			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				conn.UserLock.Lock()
				//select {
				//case conn.Send <- message:
				//	flag = true
				//default:
				//	close(conn.Send)
				//	delete(Manager.Clients, conn.ID)
				//}
				conn.Send <- message
				flag = true
				conn.UserLock.Unlock()
			}
			if flag{
				log.Println("对方在线应答")
				replyMsg := &ReplyMsg{
					Code:    200,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
			}else{
				log.Println("对方不在线")
				replyMsg := &ReplyMsg{
					Code:    200,
					Content: "对方不在线",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func WsHandler(c *gin.Context){
	Uid:=c.Query("uid")
	SendID:=c.Query("toid")
	fmt.Println("uid,toid",Uid,SendID)
	conn,err:=upGrader.Upgrade(c.Writer,c.Request,nil)

	fmt.Println("err",err)
	if err!=nil{
		http.NotFound(c.Writer,c.Request)
		return
	}
	user:=&User{
		ID:Uid+"->"+SendID,
		SendID: SendID+"->"+Uid,
		Socket: conn,
		Send:	make(chan []byte),
	}
	//server:=NewServer(Uid,SendID)
	//server.BroadCastMessage(user,"上线了")
	Manager.Register <- user
	go user.Read()
	go user.Write()

}

func (this*User)Read() {
	defer func() {
		_=this.Socket.Close()
	}()
	for{
		this.Socket.PongHandler()
		SendMsg:=new(SendMsg)
		err := this.Socket.ReadJSON(&SendMsg)
		if err != nil{
			log.Println("传入数据格式不正确", err)
			Manager.Unregister <- this
			_ = this.Socket.Close()
			break
		}
		log.Println("SendMsg.Content", SendMsg.Content)
		fmt.Println(SendMsg.Content)
		if SendMsg.Type==1{
			log.Println(this.ID, "发送消息:", SendMsg.Content)
			Manager.Broadcast <- &Broadcast{
				Client:  this,
				Message: []byte(SendMsg.Content),
			}
		}
	}
}
func (this *User)Write()  {
	defer func() {
		_=this.Socket.Close()
	}()
	for{
		select {
			case message,ok := <- this.Send :
				if !ok {
					_ = this.Socket.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				log.Println(this.ID,"接受消息:",string(message))
				replymsg:=&ReplyMsg{
					Code:    e.WebsocketSuccessMessage,
					Content: fmt.Sprintf("%s", string(message)),
				}
				msg,_:=json.Marshal(replymsg)
				this.Socket.WriteMessage(websocket.TextMessage,msg)

		}
	}
}
