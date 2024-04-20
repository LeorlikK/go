package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Server() {
	// Создаем слушатель на порту 8011
	listener, err := net.Listen("tcp", ":8011")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on port 8011...")

	// Принимаем входящие соединения от клиентов
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Запускаем обработчик для каждого клиента в отдельной горутине
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Client connected: %s\n", conn.RemoteAddr())

	// Бесконечный цикл для чтения сообщений от клиента
	scanner := bufio.NewScanner(conn)
	for {
		// Читаем очередное сообщение от клиента
		if !scanner.Scan() {
			break // Выход из цикла при ошибке чтения или закрытии соединения
		}
		message := scanner.Text()
		fmt.Printf("Received message from client %s: %s\n", conn.RemoteAddr(), message)

		go func(msg string) {
			my := fmt.Sprintf("%s...my message add", msg)
			_, err := conn.Write([]byte(my + "\n"))
			if err != nil {
				log.Printf("Error sending message to client %s: %v", conn.RemoteAddr(), err)
			}
		}(message)
	}
	fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr())
}
