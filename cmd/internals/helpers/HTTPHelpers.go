package helpers

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func HTTPHelper(eCh chan int, wg *sync.WaitGroup, link, fileName string) {
	defer wg.Done()
	var req, reqErr = http.NewRequest("GET", link, nil)

	if reqErr != nil {
		eCh <- +1
		return
	}

	var res, resErr = http.DefaultClient.Do(req)

	if resErr != nil {
		eCh <- +1
		return
	}

	var fileByte, byteErr = io.ReadAll(res.Body)

	if byteErr != nil {
		eCh <- +1
		return
	}

	if err := os.WriteFile(fileName, fileByte, 0666); err != nil {
		eCh <- +1
		return
	}
}
