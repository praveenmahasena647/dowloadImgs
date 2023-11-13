package helpers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/praveenmahasena647/imgDownload/cmd/internals/lists"
)

func WriteFile(l lists.List) error {
	var errCount = make(chan int)
	var ptr = l.Head

	for ptr != nil {
		var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Millisecond)
		go fileWrite(ctx, ptr, errCount)
		ptr = ptr.Next
		cancel()
	}
	close(errCount)
	return fmt.Errorf("failed to create %v files", <-errCount)
}

func fileWrite(ctx context.Context, n *lists.Node, c chan<- int) {
	var fileBytes = make(chan []byte)
	go HTTPHelper(ctx, n.Link, fileBytes, c)
	defer close(fileBytes)

	var err = os.WriteFile(n.Name, <-fileBytes, 0666)

	if err != nil {
		c <- +1
	}
}
