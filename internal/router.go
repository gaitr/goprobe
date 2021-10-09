package internal

import (
	"github.com/gaitr/goprobe/internal/request"
	"net/http"
)

func Router(client http.Client, args []string, flagPool *request.FlagPool) {
	url := args[0]

	var sendRequest request.Process = &request.HeadRequest{}

	if flagPool.IsGet {
		sendRequest = &request.GetRequest{}
	}

	sendRequest.PrintResponse(sendRequest.SendRequest(client, url))
}
