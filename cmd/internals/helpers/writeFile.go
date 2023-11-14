package helpers

import (
	"fmt"
	"sync"

	"github.com/praveenmahasena647/imgDownload/cmd/internals/lists"
)

func WriteFile(l lists.List) error {
	var ptr = l.Head
	var wg = &sync.WaitGroup{}
	var errCountCh = make(chan int)
	for ptr != nil {
		wg.Add(1)
		go HTTPHelper(errCountCh, wg, ptr.Link, ptr.Name)
		ptr = ptr.Next
	}
	wg.Wait()
	close(errCountCh)
	return fmt.Errorf("these many files couldn't be written %v", <-errCountCh)
}
