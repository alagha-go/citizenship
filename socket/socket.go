package socket

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

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


func OnConnection(channel *socketio.Channel) {
	if len(channel.Request().Header["Cookie"]) <= 0  {
		channel.Emit("password", `{"error": "unauthorized"`)
		time.Sleep(800*time.Millisecond)
		channel.Close()
		return
	}
	data, err := ioutil.ReadFile("./password.txt")
	if err != nil {
		channel.Emit("password", `{"error": "could not read password from our database"`)
		time.Sleep(800*time.Millisecond)
		channel.Close()
		return
	}
	password := string(data)
	cookie := channel.Request().Header["Cookie"][0]
	if password != cookie {
		channel.Emit("password", `{"error": "wrong password"`)
		time.Sleep(800*time.Millisecond)
		channel.Close()
		return
	}
}