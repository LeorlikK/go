package html

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func Main() {
	// response, err := http.Get("http://127.0.0.1:8000/tee")
	response, err := http.Get("https://p2p.binance.com/en")
	if err != nil {
		panic("Error: request error")
	}

	defer response.Body.Close()

	fmt.Println(response.Status)
	body := response.Body

	html, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error: get html error")
	}

	html.Find("table").Each(func(tableIndex int, table *goquery.Selection) {
		html, err := table.Html()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(html)
		// table.Find("tbody").Each(func(tbodyIndex int, tbody *goquery.Selection) {

		// tbody.Find("tr").Each(func(trIndex int, tr *goquery.Selection) {

		// })
		// })
	})
	fmt.Println("END FIND...")
}
