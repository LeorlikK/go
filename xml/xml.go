package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Feed struct {
	XMLName     string `xml:"Feed"`
	FeedVersion int    `xml:"Feed_Version"`
	Users       []User `xml:"Object"`
}

type User struct {
	Id     int    `xml:"id"`
	Name   string `xml:"name"`
	Family string `xml:"family"`
	Age    int    `xml:"age"`
}

func Main() {
	var users []User
	var feed Feed

	userOne := User{1, "One_Name", "One_Family", 99}
	userTwo := User{2, "Two_Name", "Two_Family", 100}
	users = append(users, userOne, userTwo)

	feed = Feed{XMLName: "Feed", FeedVersion: 3, Users: users}

	xmlData, err := xml.MarshalIndent(feed, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.Create("./xml/test.xml")
	_, err = file.Write(xmlData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("END")
}
