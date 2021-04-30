package server

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

type responseErr struct {
	Error string `json:"error"`
}

func replyError(ctx *fasthttp.RequestCtx, log *zerolog.Logger, status int, message string) {
	ctx.SetStatusCode(status)

	msg := responseErr{}
	msg.Error = message

	body, err := json.Marshal(msg)
	if err != nil {
		log.Err(err).Msg("can't marshal response body")
		return
	}

	ctx.SetBody(body)
}
