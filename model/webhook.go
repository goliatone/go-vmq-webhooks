package model

import (
	"fmt"
	"encoding/base64"
)
//TODO: missing all the _m5 endpoints

// Topic struct
type Topic struct {
	Topic string `json:"topics,omitempty"`
	QoS int `json:"qos,omitempty"`
}

//Reply struct
type Reply struct {
	//ok
	Result string `json:"result"`
	Topics []Topic `json:"topics,omitempty"`
}

//NoAuthReply should be an empty object
//on_publish
//on_subscribe
//on_offline_message
//on_client_wakeup
//on_client_offline
// on_client_gone
type NoAuthReply struct {}

//NextReply ...
type NextReply struct {
	//next
	Result string `json:"result"`
}

// ModifiersReply ...
// {
//     "result": "ok",
//     "modifiers": {
//         "topic": "rewritten/topic",
//         "qos": 2,
//         "payload": "rewritten payload",
//         "retain": true
//     }
// }
type ModifiersReply struct {
	//ok
	Result string `json:"result"`
	Modifiers Modifier `json:"modifiers,omitempty"`
}

//Modifier ...
type Modifier struct {
	Topic string `json:"topics,omitempty"`
	QoS int `json:"qos,omitempty"`
	//This is base64
	Payload interface{} `json:"payload,omitempty"`
	Retain bool `json:"retain,omitempty"`
}

//UnsubscribeReply ...
// {
//     "result": "ok",
//     "topics":
//         ["rewritten/topic"]
// }
type UnsubscribeReply struct {
	Result string `json:"result"`
	Topics []string `json:"topics,omitempty"`
}

//ErrorReply result
type ErrorReply struct {
	Result ErrorMessage `json:"result"`
}

//ErrorMessage error message
type ErrorMessage struct {
	Message string `json:"error"`
}

//OnRegister ...
// {
//     "peer_addr": "127.0.0.1",
//     "peer_port": 8888,
//     "username": "username",
//     "mountpoint": "",
//     "client_id": "clientid"
// }
type OnRegister struct {
	PeerAddress string `json:"peer_addr,omitempty"`
	PeerPort uint16 `json:"peer_port,omitempty"`
	Username string `json:"username,omitempty"`
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	CleanSession bool `json:"clean_session,omitempty"`
}

//OnSubscribe ...
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
type OnSubscribe struct {
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`	
	Username string `json:"username,omitempty"`
	Topics []Topic `json:"topics,omitempty"`
}

//OnUnsubscribe ...
// {
//     "username": "username",
//     "client_id": "clientid",
//     "mountpoint": "",
//     "topics":
//         ["a/b", "c/d"]
// }
type OnUnsubscribe struct {
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`	
	Username string `json:"username,omitempty"`
	Topics []string `json:"topics,omitempty"`
}

//Payload ...
type Payload string

//UnmarshalJSON ...
func (i *Payload) UnmarshalJSON(b []byte) error {
	
	data, err := base64.StdEncoding.DecodeString(string(b))
    if err != nil {
        return err
	}
	fmt.Printf("payload %s\n", data)

	*i = Payload(data)
	return nil
}
//OnPublish ...
// {
//     "username": "username",
//     "client_id": "clientid",
//     "mountpoint": "",
//     "qos": 1,
//     "topic": "a/b",
//     "payload": "hello",
//     "retain": false
// }
type OnPublish struct {
	Topic string `json:"topics,omitempty"`
	QoS int `json:"qos,omitempty"`
	//This is base64
	Payload interface{} `json:"payload,omitempty"`
	Retain bool `json:"retain,omitempty"`
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`	
	Username string `json:"username,omitempty"`	
}

//OnDeliver ...
type OnDeliver struct {
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`	
	Username string `json:"username,omitempty"`
	Topic string `json:"topic,omitempty"`
	//This is base64
	Payload interface{} `json:"payload,omitempty"`
}

// OnOfflineMessage ...
type OnOfflineMessage struct {
	Topic string `json:"topics,omitempty"`
	QoS int `json:"qos,omitempty"`
	//This is base64
	Payload interface{} `json:"payload,omitempty"`
	Retain bool `json:"retain,omitempty"`
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`	
}

// OnClientWakeUp ...
type OnClientWakeUp struct {
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`
}

//OnClientOffline on_client_offline
type OnClientOffline struct {
	OnClientWakeUp
}

//OnClienGone on_client_gone
type OnClienGone struct {
	OnClientWakeUp
}

// OnDeliverReply ...
// {
// 	"result": "ok",
// 	"modifiers":
// 	{
// 		  "topic": "rewritten/topic",
// 		  "payload": "rewritten payload"
// 	  }
//   }
type OnDeliverReply struct {
	Result string `json:"result"`
	Modifiers Modifier `json:"modifiers,omitempty"`
}


//AuthOnRegister is the payload for auth_on_register
// {
//     "peer_addr": "127.0.0.1",
//     "peer_port": 8888,
//     "username": "username",
//     "password": "password",
//     "mountpoint": "",
//     "client_id": "clientid",
//     "clean_session": false
// }
type AuthOnRegister struct {
	OnRegister
	Password string `json:"password,omitempty"`
}

//AuthOnSubscribe struct
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
type AuthOnSubscribe struct {
	OnSubscribe
	//TODO: Do we need this?
	Password string `json:"password,omitempty"`
}


//AuthOnPublish auth_on_publish hook payload
// {
//     "username": "username",
//     "client_id": "clientid",
//     "mountpoint": "",
//     "qos": 1,
//     "topic": "a/b",
//     "payload": "hello",
//     "retain": false
// }
type AuthOnPublish struct {
	Topic string `json:"topics,omitempty"`
	QoS int `json:"qos,omitempty"`
	//This is base64
	Payload interface{} `json:"payload,omitempty"`
	Retain bool `json:"retain,omitempty"`
	Mountpoint string `json:"mountpoint,omitempty"`
	ClientID string `json:"client_id,omitempty"`	
	Username string `json:"username,omitempty"`
}
