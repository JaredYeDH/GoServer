package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"time"
)

var url = "ws://192.168.1.106:9000/"
var origin = "http://192.168.1.106/"

type MsgHead struct {
	SeqID      int    `json:"seq"`    //! 顺序标志
	MsgType    string `json:"type"`   //! 消息类型
	Action     string `json:"action"` //! 消息行为
	CreateTime int64  `json:"time"`   //! 生成时间
}

type PlayerLogin struct {
	MsgHead   `json:"head"`
	AccountID int64
	LoginKey  string
}

//! websocket 连接测试
func TestSendMsg() {
	//! 使用默认配置连接服务端
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println("connect server fail!")
		return
	}

	msg := new(PlayerLogin)
	msg.SeqID = 1
	msg.MsgType = "login"
	msg.Action = "login"
	msg.CreateTime = time.Now().Unix()
	msg.AccountID = 1
	msg.LoginKey = "566a3412aeddbf1eb837574e"

	b, _ := json.Marshal(msg)

	fmt.Println("send: ", string(b))

	err = websocket.JSON.Send(ws, msg)
	if err != nil {
		fmt.Println("send to server fail!")
		return
	}

	// fmt.Println("send msg: ", "Hello~~")

	// recv := ""
	// websocket.JSON.Receive(ws, &recv)

	// fmt.Println("recv msg: ", recv)

	// fmt.Println("Success!")
	defer ws.Close()
}

//! 创建角色信息
type Msg_CreateRole struct {
	MsgHead    `json:"head"`
	AccountID  int64  `json:"accountid"`
	LoginKey   string `json:"key"`
	PlayerName string `json:"name"`
}

func TestCreateRole() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println("connect server fail!")
		return
	}

	msg := Msg_CreateRole{}
	msg.SeqID = 1
	msg.MsgType = "login"
	msg.Action = "create"
	msg.CreateTime = time.Now().Unix()
	msg.AccountID = 1
	msg.LoginKey = "566a9092aeddbf24a412f134"
	msg.PlayerName = "Arteezy"

	err = websocket.JSON.Send(ws, &msg)
	if err != nil {
		fmt.Println("send msg fail")
		return
	}

	b, _ := json.Marshal(msg)

	fmt.Println("send: ", string(b))

	defer ws.Close()
}
