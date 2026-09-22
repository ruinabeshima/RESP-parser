package main 

import (
	"fmt"
	"bufio"
	"os"
)

// Data types correspond to symbol of first byte
const (
	simpleString = "+"
	simpleError = "-"
	integer = ":"
	bulkString = "$"
	array = "*"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter RESP command: ")
	if scanner.Scan() {
		input := scanner.Text() 
		fmt.Println("Your command is:", input)
	}
}