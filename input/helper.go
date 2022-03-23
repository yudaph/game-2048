package input

import (
	"bufio"
	"os"
	"strings"
)

func input(isMove bool) (str string, err error) {
	reader := bufio.NewReader(os.Stdin)
	str, err = reader.ReadString('\n')
	if err != nil {
		return
	}

	str = trimString(str)
	if isMove {
		if len(str) > 0 {
			str = str[:1]
		}
	}

	return
}

func trimString(str string) string {
	return strings.TrimSpace(strings.TrimSuffix(str, "\n"))
}
