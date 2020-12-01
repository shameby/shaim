package server

import (
	"fmt"
	"encoding/json"

	"github.com/googollee/go-socket.io"
	"shaim/check"
)

var nameM map[string]string
var nameId map[string]string
var conM map[string]socketio.Conn

type say struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var S *socketio.Server

func init() {
	nameM = make(map[string]string)
	nameId = make(map[string]string)
	conM = make(map[string]socketio.Conn)

	S, _ = socketio.NewServer(nil)

	S.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		if _, exist := conM[s.ID()]; !exist {
			conM[s.ID()] = s
		}
		return nil
	})

	S.OnEvent("/", "join", func(s socketio.Conn, username string) {
		s.Join("class50")
		if _, exist := nameM[s.ID()]; !exist {
			nameM[s.ID()] = username
			nameId[username] = s.ID()
		}
		check.SetHash("class50", username)
		s.Emit("msg", username+",您已加入房间：class50")
	})

	S.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		grpcBroadcasting("/", "class50", nameM[s.ID()]+": "+msg)
		return "发送成功：" + msg
	})

	S.OnEvent("/", "say", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		data := &say{}
		json.Unmarshal([]byte(msg), data)
		s.Emit("msg", "你 对 "+data.Name +"说: "+data.Msg)
		if sId, exist := nameId[data.Name]; !exist {
			err := grpcSay("/", "class50", data.Name, nameM[s.ID()]+" 对你说: "+data.Msg)
			if err != nil {
				s.Emit("msg", err.Error())
				return err.Error()
			}
		} else {
			if toConn, exist := conM[sId]; !exist {
				err := grpcSay("/", "class50", data.Name, nameM[s.ID()]+" 对你说: "+data.Msg)
				if err != nil {
					s.Emit("msg", err.Error())
					return err.Error()
				}
			} else {
				toConn.Emit("msg", nameM[s.ID()]+" 对你说: "+data.Msg)
			}
		}

		return "发送成功：" + msg
	})

	S.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	S.OnDisconnect("/", func(s socketio.Conn, msg string) {
		check.DelHash("class50", nameM[s.ID()])
		fmt.Println(s.ID()+" closed", msg)
	})
}
