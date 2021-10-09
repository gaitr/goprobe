package request

import (
	"fmt"
	"net/http"
	"time"
)

type HeadResponse struct {
	statusCode    int
	contentLength int64
	lastModified  string
}

func Head(client http.Client, path string) HeadResponse {
	response, err := client.Head(path)
	headResponse := HeadResponse{}

	if err == nil {
		headResponse.statusCode = response.StatusCode

		conLength := response.Header.Get("Content-Length")
		if conLength != "" {
			headResponse.contentLength = response.ContentLength
		}

		t, err := time.Parse(time.RFC1123, response.Header.Get("Last-Modified"))
		if err == nil {
			headResponse.lastModified = t.String()
		}
	}
	return headResponse
}

func CompleteHeadResponse(headResponse HeadResponse) {
	statusCode := headResponse.statusCode
	contentLength := headResponse.contentLength
	lastModified := headResponse.lastModified
	result := "Please provide a valid URL"

	if statusCode != 0 {
		result = fmt.Sprintf("Status Code: %d ", headResponse.statusCode)
		if contentLength != 0 {
			result += fmt.Sprintf("Content-Length: %d ", contentLength)
		}

		if lastModified != "" {
			result += fmt.Sprintf("Last-Modified: %s", lastModified)
		}
	}
	fmt.Println(result)
}

func DefaultResponse(headResponse HeadResponse, path string) {
	statusCode := headResponse.statusCode
	result := "Please provide a valid URL"

	if statusCode != 0 {
		result = fmt.Sprintf("Status Code: %d ", headResponse.statusCode)
		if path != "" {
			result += fmt.Sprintf(" URL: %s", path)
		}
	}
	fmt.Println(result)
}
