package socket

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func Client() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "127.0.0.1:8011")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			fmt.Println("Received message from server:", message)
		}
	}()

	// Отправляем сообщения на сервер каждые 5 секунд
	for {
		message := []string{
			"https://jsonplaceholder.typicode.com/posts/1",
			"https://jsonplaceholder.typicode.com/posts/2",
			"https://jsonplaceholder.typicode.com/posts/3",
		}
		stringMessage, _ := json.Marshal(message)
		_, err := conn.Write(append(stringMessage, '\n'))
		if err != nil {
			log.Println("Error writing to server:", err)
			break
		}
		fmt.Println("Sent message to server:", message)
		time.Sleep(5 * time.Second)
	}
}
