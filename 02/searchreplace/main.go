package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	input := os.Args[1]
	search := os.Args[2]
	replacement := os.Args[3]

	var result strings.Builder
	for _, r := range input {
		if string(r) == search {
			result.WriteString(replacement)
		} else {
			result.WriteRune(r)
		}
	}

	fmt.Println(result.String())
}
