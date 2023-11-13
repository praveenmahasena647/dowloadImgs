package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/imgDownload/cmd/internals/helpers"
)

func main() {
	var links = helpers.GetLinks()
	var err error = helpers.WriteFile(links)

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
