package internal

import (
	"fmt"
	"github.com/gaitr/goprobe/internal/request"
	"net/http"
	"os"
)

func Router(client http.Client, path string, flagPool *request.FlagPool) error {
	var sendRequest request.Process = &request.HeadRequest{}

	if flagPool.IsGet {
		sendRequest = &request.GetRequest{}
	}

	_, e := fmt.Fprintln(
		os.Stdout, sendRequest.PrintResponse(sendRequest.SendRequest(client, path)))

	if e != nil {
		return e
	}
	return nil
}
