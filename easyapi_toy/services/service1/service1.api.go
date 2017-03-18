package service1

import (
	"easyapi"
	"easyapi/easyapi_toy/services"
	"net"
)

type MessageID byte

const (
	MsgID_Add MessageID = 2
)

func (this *AddIn) ServiceID() byte {
	return byte(services.ServiceID_Service1)
}

func (this *AddIn) MessageID() byte {
	return byte(MsgID_Add)
}

func (this *AddIn) Identity() string {
	return ""
}

func (this *AddOut) ServiceID() byte {
	return byte(services.ServiceID_Service1)
}

func (this *AddOut) MessageID() byte {
	return byte(MsgID_Add)
}

func (this *AddOut) Identity() string {
	return ""
}

func (this *Service1) NewRequest(msgID byte) easyapi.IMessage {
	switch MessageID(msgID) {
	case MsgID_Add:
		return &AddIn{}
	}
	return nil
}

func (this *Service1) HandleRequest(conn net.Conn, msg easyapi.IMessage) {
	switch MessageID(msg.MessageID()) {
	case MsgID_Add:
		this.Add(conn, msg.(*AddIn))
	}
}

func (this *Service1) ServiceID() byte {
	return byte(services.ServiceID_Service1)
}
