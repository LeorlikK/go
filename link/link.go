package link

import "fmt"

func Link() {
	linkOne := "TEST_LINK_STRING_ONE"
	linkTwo := "TEST_LINK_STRING_TWO"
	PrintStringDefault(linkOne)
	PrintStringLink(&linkTwo)
	fmt.Println(linkOne)
	fmt.Println(linkTwo)
	fmt.Println(&linkTwo)
}

func PrintStringDefault(link string) {
	link += "_ADD_MESSAGE_DEFAULT"
	fmt.Println(link)
}

func PrintStringLink(link *string) {
	*link += "_ADD_MESSAGE_LINK"
	fmt.Println(link)
	fmt.Println(&link)
}
