package lib

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//MasterSocketConnection contains reference to the main websocket connection
var MasterSocketConnection *websocket.Conn

//SetupSockServer inits the master communication socket
func SetupSockServer() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error
		MasterSocketConnection, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		go PingClient()
	})
	fmt.Println("starting sockserver")
	http.ListenAndServe(":9000", nil)
}

//PingClient periodically pings the connected clients
//TODO
func PingClient() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if MasterSocketConnection != nil {
				if err := MasterSocketConnection.WriteMessage(1, []byte("PING")); err != nil {
					return
				}
			}
		}
	}
}
