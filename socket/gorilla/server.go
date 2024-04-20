package gorilla

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	//defer conn.Close()
	fmt.Println("111")
	go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		go func(msg []byte) {
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		}(p)
	}
}

func handleUpgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// обработка ошибки
		return
	}
	defer conn.Close()
	// дальнейшая обработка соединения
}

func Server() {
	http.HandleFunc("/ws", handler)
	log.Println("http server started on :8012")
	err := http.ListenAndServe(":8012", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Client() {

}
