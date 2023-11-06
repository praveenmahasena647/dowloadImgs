package links

import (
	"io"
	"net/http"
	"os"
)

func GetImg(links map[string]string) error {

	for name, link := range links {
		var con, conErr = http.Get(link)
		if conErr != nil {
			return conErr
		}
		defer con.Body.Close()
		var content, contentErr = io.ReadAll(con.Body)
		if contentErr != nil {
			return contentErr
		}
		var fileErr = os.WriteFile(name, content, 0666)

		if fileErr != nil {
			return fileErr
		}
	}

	return nil
}
