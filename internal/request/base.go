package request

import (
	"fmt"
	res "github.com/gaitr/goprobe/internal/response"
	"net/http"
)

type Process interface {
	SendRequest(client http.Client, path string) res.Response
	PrintResponse(response res.Response) string
}

type FlagPool struct {
	IsGet bool
}

type Request struct {
	Type string
}

func (r *Request) PrintResponse(response res.Response) string {
	statusCode := response.StatusCode
	contentLength := response.ContentLength
	lastModified := response.LastModified
	result := "Please provide a valid URL"

	if statusCode != 0 {
		result = fmt.Sprintf("Status Code: %d ", response.StatusCode)
		if contentLength != 0 {
			result += fmt.Sprintf("Content-Length: %d ", contentLength)
		}

		if lastModified != "" {
			result += fmt.Sprintf("Last-Modified: %s", lastModified)
		}
	}
	return r.Type + " " + result + " " + response.Path
}
