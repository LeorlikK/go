package file

import (
	"fmt"
	"log"
	"os"
)

const path = "./file/test.txt"

func Main() {
	// readFile()
	// writeFile()
	appendFile()
}

func readFile() {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func writeFile() {
	data := []byte("Hello Four")

	if err := os.WriteFile(path, data, 0777); err != nil {
		log.Fatal(err)
	}

	fmt.Println("write")
}

func appendFile() {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err = file.WriteString("Hello Six"); err != nil {
		log.Fatal(err)
	}
}
