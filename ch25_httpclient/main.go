package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	/*clients := make([]http.Client, 3)
	for index, client := range clients {

		jar, err := cookiejar.New(nil)
		if err == nil {
			client.Jar = jar
		}
		for i := 0; i < 3; i++ {
			req, err := http.NewRequest(http.MethodGet,
				"http://localhost:5000/cookie", nil)
			if err == nil {
				response, err := client.Do(req)
				if err == nil && response.StatusCode == http.StatusOK {
					fmt.Fprintf(os.Stdout, "Client %v: ", index)
					io.Copy(os.Stdout, response.Body)
					defer response.Body.Close()
				} else {
					Printfln("Request Error: %v", err.Error())
				}
			} else {
				Printfln("Request Init Error: %v", err.Error())
			}
		}
	}*/

	//redirects
	http.DefaultClient.CheckRedirect = func(req *http.Request,
		previous []*http.Request) error {
		if len(previous) == 3 {
			url, _ := url.Parse("http://localhost:5000/html")
			req.URL = url
		}
		return nil
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/redirect1", nil)
	if err == nil {
		var response *http.Response
		response, err = http.DefaultClient.Do(req)
		if err == nil {
			io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

}
