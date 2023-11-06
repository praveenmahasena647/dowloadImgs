package links

import (
	"bufio"
	"os"
)

func GetLinks() map[string]string {
	var links map[string]string = map[string]string{}
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var keyLink string = scanner.Text()
		if keyLink != "end" {
			var name, link string = getVals(keyLink)
			links[name] = link
			continue
		}
		break
	}
	return links
}

func getVals(keyLink string) (string, string) {
	var key, link, leng = []byte{}, "", len(keyLink)

	for i := 0; i < leng; i++ {
		if keyLink[i] == ':' {
			link = keyLink[i+1:]
			break
		}
		key = append(key, keyLink[i])
	}
	return string(key), link
}
