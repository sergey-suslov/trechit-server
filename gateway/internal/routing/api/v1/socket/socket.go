package socket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// InitSocketConn Init socket connection for new connection
func InitSocketConn(e echo.Context) error {
	conn, err := upgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		log.Println("Error creating socket connection", err)
		return err
	}
	defer conn.Close()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading socket message", err)
			return err
		}
		log.Println("Got message", msgType, msg)
		// TODO: implement message handling
	}
}
