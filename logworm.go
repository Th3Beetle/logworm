package logworm

import (
	"strconv"
	"strings"
)

type Request struct {
	id            int
	method        string
	uri           string
	statusCode    string
	contentLength string
	request       string
	response      string
}

func (r *Request) getLogSlice() []interface{} {
	res := []interface{}{}
	res = append(res, r.id, r.method, r.uri, r.statusCode, r.contentLength)
	return res
}

func (r *Request) parse() {
	lines := strings.Split(r.request, "\r\n")
	first_line := strings.Split(lines[0], " ")
	r.method = first_line[0]
	r.uri = first_line[1]

	r.statusCode = strings.Split(r.response, " ")[1]
	r.contentLength = strings.Split(strings.Split(r.response, "Content-Length: ")[1], "\r\n")[0]
}

func StartLog(requests chan string, response chan string, window_slices chan []interface{}, request_id chan string, window_requests chan string, window_responses chan string) {
	requestLog := make(map[int]Request)
	var counter int
	go listenForQuery(&requestLog, request_id, window_requests, window_responses)
	for {

		newRequest := <-requests
		newResponse := <-response
		var newParsedRequest Request
		newParsedRequest.request = newRequest
		newParsedRequest.response = newResponse
		newParsedRequest.id = counter
		newParsedRequest.parse()
		requestLog[counter] = newParsedRequest
		window_slices <- newParsedRequest.getLogSlice()
		counter += 1
	}
}

func listenForQuery(requestLog *map[int]Request, request_id chan string, window_requests chan string, window_responses chan string) {
	stringId := <-request_id
	id, _ := strconv.Atoi(stringId)
	request := (*requestLog)[id]
	window_requests <- request.request
	window_responses <- request.response
}
