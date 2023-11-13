package helpers

import (
	"bufio"
	"os"

	"github.com/praveenmahasena647/imgDownload/cmd/internals/lists"
)

func GetLinks() lists.List {
	var list = lists.List{}
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var link string = scanner.Text()
		if link != "end" {
			var node = makeNode(link)
			list.Add(node)
			continue
		}
		break
	}
	return list
}

func makeNode(link string) lists.Node {
	var node = lists.Node{}
	var leng = len(link)

	for i := 0; i < leng; i++ {
		if link[i] != ':' {
			node.Name = link[:i+1]
			continue
		}
		node.Link = link[i+1:]
		break
	}
	return node
}
