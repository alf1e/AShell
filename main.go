package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		cmd := exec.Command(fmt.Sprintf("/usr/bin/%s", text))
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Print(string(stdout))
		}
	}
}

func parseText(textToParse string) {
	if textToParse == "exit" {
		os.Exit(0)
	}
}
