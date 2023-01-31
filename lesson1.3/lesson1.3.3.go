package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var abbr string

	phrase := readString()
	words := strings.Fields(phrase)

	for _, word := range words {
		firstCharacter := []rune(word)[0:1]

		if unicode.IsLetter(firstCharacter[0]) {
			abbr += string(unicode.ToUpper(firstCharacter[0]))
		}
	}

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
