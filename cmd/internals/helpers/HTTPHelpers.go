package helpers

import (
	"context"
	"io"
	"net/http"
)

func HTTPHelper(ctx context.Context, link string, b chan<- []byte, c chan<- int) {
	var req, reqErr = http.NewRequestWithContext(ctx, "GET", link, nil)
	if reqErr != nil {
		c <- +1
		return
	}
	var res, resErr = http.DefaultClient.Do(req)
	if resErr != nil {
		c <- +1
		return
	}
	defer res.Body.Close()
	var bytes, byteErr = io.ReadAll(res.Body)
	if byteErr != nil {
		c <- +1
		return
	}
	b <- bytes
}
