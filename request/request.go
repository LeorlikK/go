package request

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var urlsSlice = []string{
	"https://httpbin.org/headers",
	"https://api.ipify.org/?format=json",
}

func Request() {
	typeRequest := "RoundTrip"
	client := http.Client{}
	request, err := http.NewRequest("GET", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, urlString := range urlsSlice {
		urlParse, err := url.Parse(urlString)
		request.URL = urlParse

		if typeRequest == "CLIENT" {
			proxyURL, err := url.Parse("http://hjrKNO:FlWcsJkvDX@188.130.188.166:5500")
			if err != nil {
				log.Fatal(err)
			}
			// header
			request.Header.Set("accept", "application/json, text/plain, */*")
			request.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
			// cookie
			cookieURL, err := url.Parse(urlString)
			cookies := []*http.Cookie{
				{Name: "cookie_name_1", Value: "cookie_value_1"},
				{Name: "cookie_name_2", Value: "cookie_value_2"},
			}
			jar, err := cookiejar.New(nil)
			if err != nil {
				log.Fatal(err)
			}
			jar.SetCookies(cookieURL, cookies)
			client.Jar = jar
			// proxy
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
			client.Transport = transport

		} else if typeRequest == "REQUEST" {
			// header
			request.Header.Set("HeaderName1", "HeaderValue1")
			request.Header.Set("HeaderName2", "HeaderValue2")
			// cookie
			cookies := []*http.Cookie{
				{Name: "cookie_name_1", Value: "cookie_value_1"},
				{Name: "cookie_name_2", Value: "cookie_value_2"},
			}
			for _, cookie := range cookies {
				request.AddCookie(cookie)
			}
			// proxy
			proxyURL, err := url.Parse("http://hjrKNO:FlWcsJkvDX@188.130.188.166:5500")
			if err != nil {
				log.Fatal(err)
			}
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
			client.Transport = transport
		} else if typeRequest == "RoundTrip" {
			proxyURL, err := url.Parse("http://hjrKNO:FlWcsJkvDX@188.130.188.166:5500")
			if err != nil {
				log.Fatal(err)
			}
			transportDefault := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
			transport := &myTransportWithHeaders{
				headers: map[string]string{
					"accept":          "application/json, text/plain, */*",
					"accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
				},
				transport: transportDefault, // http.DefaultTransport
			}
			client.Transport = transport
		}

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}
}

type myTransportWithHeaders struct {
	headers   map[string]string
	transport http.RoundTripper
}

func (t *myTransportWithHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}

	return t.transport.RoundTrip(req)
}
