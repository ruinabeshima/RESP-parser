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

func is_CRLF(str string, pointer int) bool {

	// Pointer out of bounds
	if pointer + 1 >= len(str) || pointer < 0 {
		return false
	}

	// Check if there is a Carriage Return Line Feed (\r\n)
	if str[pointer] == '\r' && str[pointer + 1] == '\n' {
		return true 
	}

	return false 
}