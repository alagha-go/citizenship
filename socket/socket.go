package socket

import (
	"log"
	"net/http"

	socketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
)

var (
	PORT = ":12000"
	Server = socketio.NewServer(transport.GetDefaultWebsocketTransport())
)

func StartSocketServer() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", Server)

	log.Println("Starting SocketServer...")
	log.Panic(http.ListenAndServe(PORT, serveMux))
}