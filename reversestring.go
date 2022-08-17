package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(statement string) (reverse string) {
	for _, i := range statement {
		reverse = string(i) + reverse
	}
	return
}

func main() {
	var statement string
	fmt.Println("Hello")
	fmt.Print("Please enter a statement to reverse: ")
	reader := bufio.NewReader(os.Stdin)
	statement, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred, please try again.")
		return
	}
	statement = strings.TrimSuffix(statement, "\n")
	strRev := reverse(statement)
	fmt.Print(strRev)
}
