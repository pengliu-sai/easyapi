package service1

import (
	"easyapi"
	"fmt"
	"net"

	"github.com/funny/binary"
)

type Service1 struct {
}

func (this *Service1) APIs() easyapi.APIs {
	return easyapi.APIs{
		0: {AddIn{}, AddOut{}},
	}
}

type AddIn struct {
	A int
	B int
}

type AddOut struct {
	C int
}

//
// AddIn
//
func (s *Service1) Add(conn net.Conn, in *AddIn) {
	fmt.Printf("Add A: %d, B: %d\n", in.A, in.B)
	addOut := &AddOut{
		C: in.A + in.B,
	}

	var buff = binary.Buffer{Data: make([]byte, 6+addOut.Size())}
	buff.WriteUint32LE(uint32(addOut.Size()))
	buff.WriteUint8(addOut.ServiceID())
	buff.WriteUint8(addOut.MessageID())
	addOut.Marshal(buff.Data[6:])
	conn.Write(buff.Data)
}
