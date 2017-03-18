package easyapi

import (
	"log"
	"net"
	"strings"
	"time"
)

type APIs map[byte][2]interface{}

const packetHeadSize = 4 + 2

type App struct {
	services []IService
	handle   IHandle
	headBuf  []byte
	headData [packetHeadSize]byte
	listener net.Listener
}

type IService interface {
	ServiceID() byte
	NewRequest(byte) IMessage
	HandleRequest(net.Conn, IMessage)
}

type IMessage interface {
	ServiceID() byte
	MessageID() byte
	Identity() string
	Size() int
	Marshal([]byte) int
	Unmarshal([]byte) int
}

type IHandle interface {
	InitConn(net.Conn) error
	Transaction(net.Conn, *App)
}

func New() *App {
	return &App{
		services: make([]IService, 256),
		handle:   &DefaultTcpHandle{},
	}
}

func (this *App) RegisterService(s IService) {
	this.services[s.ServiceID()] = s
}

func (this *App) Listen(network, address string) (net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	this.listener = listener
	return listener, nil
}

func (this *App) Close() {
	if this.listener != nil {
		this.listener.Close()
	}
}

func (this *App) Serve(listener net.Listener) {
	log.Printf("easyapi serve on %s\n", listener.Addr().String())

	var tempDelay time.Duration

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			return
		}
		go this.handleSession(conn, this.handle)
	}
}

func (this *App) handleSession(conn net.Conn, handle IHandle) {
	handle.Transaction(conn, this)
}

func (this *App) Dial(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
