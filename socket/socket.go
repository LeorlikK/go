package socket

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

func Main() {
	serverAddress := "ws://77.239.252.141:25116"
	channelName := "newItems"

	conn, _, err := websocket.DefaultDialer.Dial(serverAddress, nil)
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}

	defer conn.Close()

	subscribeMessage := fmt.Sprintf(`{"event":"subscribe","channel":"%s"}`, channelName)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(subscribeMessage)); err != nil {
		fmt.Println("Ошибка отправки сообщения:", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Ошибка чтения сообщения:", err)
			return
		}

		fmt.Printf("Получено сообщение: %s\n", message)
	}
}

func Main2() {

	http.HandleFunc("ws://77.239.252.141:25116", handleConnections)
	fmt.Printf("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// обновление соединения до WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// цикл обработки сообщений
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("Received: %s", message)

		// эхо ансвер
		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Fatal(err)
			break
		}
	}
}

func Main3() {

}

func websocketHeandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Fatal(err)
			break
		}

		log.Printf("Req: %s", message)
	}
}
