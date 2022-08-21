package main

import (
	"fmt"

	"github.com/Th3Beetle/logworm"
)

func main() {
	requests := make(chan string)
	responses := make(chan string)
	slices := make(chan []interface{})
	id := make(chan string)
	window_requests := make(chan string)
	window_responses := make(chan string)

	go logworm.StartLog(requests, responses, slices, id, window_requests, window_responses)

	request := "GET /hello.htm HTTP/1.1\r\nUser-Agent: Mozilla/4.0 (compatible; MSIE5.01; Windows NT)\r\nHost: www.tutorialspoint.com\r\nAccept-Language: en-us\r\nAccept-Encoding: gzip, deflate\r\nConnection: Keep-Alive"
	response := "HTTP/1.1 404 Not Found\r\nDate: Sun, 18 Oct 2012 10:36:20 GMT\r\nServer: Apache/2.2.14 (Win32)\r\nContent-Length: 230\r\nConnection: Closed\r\nContent-Type: text/html; charset=iso-8859-1"
	id_val := "1"
	requests <- request
	responses <- response
	slice := <-slices
	fmt.Println(slice)
	id <- id_val
	window_request := <-window_requests
	fmt.Println(window_request)
	window_response := <-window_responses
	fmt.Println(window_response)
}
