package main

import (
	"bufio"
	"os"
	"strings"
)

func countWords(input string) map[string]int {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		freq[word]++
	}

	return freq
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var text string

	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	freq := countWords(text)

	for word, count := range freq {
		println(word+":", count)
	}
}
