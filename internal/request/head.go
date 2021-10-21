package request

import (
	res "github.com/gaitr/goprobe/internal/response"
	"net/http"
	"time"
)

type HeadRequest struct {
	request Request
}

func (hr *HeadRequest) SendRequest(client http.Client, path string) res.Response {
	response, err := client.Head(path)
	headResponse := res.Response{}
	headResponse.Path = path
	if err == nil {
		headResponse.StatusCode = response.StatusCode

		conLength := response.Header.Get("Content-Length")
		if conLength != "" {
			headResponse.ContentLength = response.ContentLength
		}

		t, err := time.Parse(time.RFC1123, response.Header.Get("Last-Modified"))
		if err == nil {
			headResponse.LastModified = t.String()
		}
	}

	return headResponse
}

func (hr *HeadRequest) PrintResponse(client http.Client, path string) string {
	response := hr.SendRequest(client, path)
	return hr.request.PrintResponse(response)
}
