package controller

import (
	"fmt"
	"encoding/base64"
	"github.com/goliatone/go-vmq-webhooks/model"

	"github.com/gin-gonic/gin"
)

//AuthOnRegister handles the `auth_on_register` hook
//It has 3 possible responses:
//{
//    "result": "ok"
//}
//
//It's also possible to override client specifics:
//{
//   "result": "ok",
//   "modifiers": {
//       "max_message_size": 65535,
//       "max_inflight_messages": 10000,
//       "retry_interval": 20000
//   }
//This will execute the next hook:
//{
// 	"result": "next"
//}
//This will reject the auth:
// {
// 	"result": {
// 	  "error": "not_allowed"
// 	}
//}
func AuthOnRegister(c *gin.Context) {
	fmt.Println("Auth on register")
	ok := model.Reply{
		Result:"ok",
	}
	c.JSON(200, ok)
}

//AuthOnPublish .
func AuthOnPublish(c *gin.Context) {
	fmt.Println("auth on publish")
	ok := model.Reply{
		Result:"ok",
	}
	c.JSON(200, ok)
}

//AuthOnSubscribe handles `auth_on_subscribe` hook
//Sample payload:
// {
//     "client_id": "clientid",
//     "mountpoint": "",
//     "username": "username",
//     "topics":
//         [{"topic": "a/b",
//           "qos": 1},
//          {"topic": "c/d",
//           "qos": 2}]
// }
//You can modify the topics to be subscribed:
// {
//     "result": "ok",
//     "topics":
//         [{"topic": "rewritten/topic",
//           "qos": 0}]
// }
func AuthOnSubscribe(c *gin.Context) {
	fmt.Println("auth on subscrbe")
	ok := model.Reply{
		Result:"ok",
	}
	c.JSON(200, ok)
}

/////////////////////////////////////

//OnPublish .
func OnPublish(c *gin.Context) {
	fmt.Println("on publish")
	var m model.OnPublish
	c.BindJSON(&m)

	fmt.Printf("client id: %s username: %s mountpoint: %s topic: %s payload: %v\n",
	m.ClientID ,
	m.Username ,
	m.Mountpoint ,
	m.Topic ,
	m.Payload)
	
	s := m.Payload.(string)
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Printf("error decoding: %s", err)
	}

	fmt.Printf("payload: %s\n", data)

	ok := model.Reply{
		Result:"ok",
	}
	c.JSON(200, ok)
}

//OnUnsubscribe .
func OnUnsubscribe(c *gin.Context) {
	fmt.Println("on unsubscribe")
	ok := model.UnsubscribeReply{
		Result: "ok",
	}
	c.JSON(200, ok)
}
//OnClientGone .
func OnClientGone(c *gin.Context) {
	fmt.Println("on client gone")
	ok := model.NoAuthReply{}
	c.JSON(200, ok)
}

//OnClientOffline .
func OnClientOffline(c *gin.Context) {
	fmt.Println("on client offline")
	var m model.OnClientOffline
	c.BindJSON(&m)
	fmt.Printf("%v\n", m)

	ok := model.NoAuthReply{}
	c.JSON(200, ok)
}
