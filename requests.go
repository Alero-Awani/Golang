package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main(){
	client := http.DefaultClient
	req, err := http.NewRequest("GET","https://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create request")
	}
	resp, err := client.Do(req)


	// resp, err := http.Post("https://httpbin.org/post", "text/plain", 
	// 	strings.NewReader("This is the request content"))
	// resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to get response from the server")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	fmt.Println(string(content))

}