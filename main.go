package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" AShell ")
	fmt.Println("--------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		parseText(text)
		fmt.Println(text)
	}
}

func parseText(textToParse string) {
	if textToParse == "exit" {
		os.Exit(0)
	}
}
