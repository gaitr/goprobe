package request

import (
	"fmt"
	res "github.com/gaitr/goprobe/internal/response"
	"net/http"
	"strconv"
)

type Process interface {
	SendRequest(client http.Client, path string) res.Response
	PrintResponse(client http.Client, path string) string
}

type FlagPool struct {
	IsGet        bool
	VerboseLevel int
}

type Request struct{}

func (r *Request) PrintResponse(response res.Response) string {
	statusCode := response.StatusCode
	contentLength := response.ContentLength
	lastModified := response.LastModified
	result := "Please provide a valid URL"

	if statusCode != 0 {
		result = strconv.Itoa(statusCode)
		if contentLength != 0 {
			result += fmt.Sprintf(" Content-Length: %d ", contentLength)
		}

		if lastModified != "" {
			result += fmt.Sprintf(" Last-Modified: %s", lastModified)
		}
	}

	return result + " " + response.Path
}
