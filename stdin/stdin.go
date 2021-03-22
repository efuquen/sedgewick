package stdin

import (
	"bufio"
	"os"
)

func ReadAllStrings() []string {
	var words []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}
