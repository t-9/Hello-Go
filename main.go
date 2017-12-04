package main

import (
	"fmt"
	"strings"
)

func main() {
	message := []string{
		"おはよしこ！",
		"だからヨハネよ！",
		"ぴょーん",
	}

	fmt.Println(strings.Join(message, "\n"))
}
