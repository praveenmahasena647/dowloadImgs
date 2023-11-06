package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/downloader/cmd/internals/links"
)

func main() {
	var link map[string]string = links.GetLinks()
	var err error = links.GetImg(link)

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}
