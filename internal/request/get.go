package request

import (
	res "github.com/gaitr/goprobe/internal/response"
	"net/http"
	"time"
)

type GetRequest struct {
	request Request
}

func (hr *GetRequest) SendRequest(client http.Client, path string) res.Response {
	response, err := client.Get(path)
	getResponse := res.Response{}
	getResponse.Path = path
	if err == nil {
		getResponse.StatusCode = response.StatusCode

		conLength := response.Header.Get("Content-Length")
		if conLength != "" {
			getResponse.ContentLength = response.ContentLength
		}

		t, err := time.Parse(time.RFC1123, response.Header.Get("Last-Modified"))
		if err == nil {
			getResponse.LastModified = t.String()
		}
	}

	return getResponse
}

func (hr *GetRequest) PrintResponse(client http.Client, path string) string {
	response := hr.SendRequest(client, path)
	return hr.request.PrintResponse(response)
}
