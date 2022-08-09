package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func get() {
	resp, err := http.Get("http://127.0.0.1:5678/abc")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	}
}

func post() {
	requestBody := "abcdefg"
	reader := strings.NewReader(requestBody)
	if resp, err := http.Post("http://127.0.0.1:5678/abc", "text/plain", reader); err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		fmt.Println(resp.Proto)
		fmt.Println(resp.Status)
		fmt.Println(resp.StatusCode)
		for k, v := range resp.Header {
			fmt.Println(k, v[0])
		}
		for _, cookie := range resp.Cookies() {
			fmt.Println(cookie.Name, cookie.Value)
		}
		io.Copy(os.Stdout, resp.Body)
	}
}

//func generalRequest() {
//	requestBode := "abcdefg"
//	reader := strings.NewReader(requestBode)
//	if request, err := http.NewRequest("POST", "http://127.0.0.1:567/abc", reader); err != nil {
//		fmt.Println(err)
//	} else {
//		request.Header.Add("User-Agent", "1234")
//		request.Header.Add("my_key", "my_value")
//		request.AddCookie(&http.Cookie{
//			Name: "skin",
//			Value: "new",
//		})
//		client:=http.Client{}
//		if resp,err:=client.Do(request);err!=nil{
//			fmt.Println(cookie.)
//		}
//	}
//}

func main() {
	//get()
	post()
}
