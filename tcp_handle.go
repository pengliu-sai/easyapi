package easyapi

import (
	"io"
	"net"

	"github.com/funny/binary"
)

type DefaultTcpHandle struct {
}

func (*DefaultTcpHandle) InitConn(conn net.Conn) error {
	return nil
}

func (*DefaultTcpHandle) Transaction(conn net.Conn, app *App) {
	app.headBuf = app.headData[:]

	if _, err := io.ReadFull(conn, app.headBuf); err != nil {
		return
	}

	packetSize := int(binary.GetUint32LE(app.headBuf[0:4]))
	packet := make([]byte, packetSize)

	if _, err := io.ReadFull(conn, packet); err != nil {
		panic(err)
	}

	serviceID := app.headBuf[4]
	msgID := app.headBuf[5]

	service := app.services[serviceID]
	req := service.NewRequest(msgID)
	req.Unmarshal(packet)
	service.HandleRequest(conn, req)
}
