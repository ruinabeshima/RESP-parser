package main

import "fmt"

// Data types correspond to symbol of first byte
const (
	simpleString = '+'
	simpleError  = '-'
	integer      = ':'
	bulkString   = '$'
	array        = '*'
)

func main() {
	testInput := "+OK\r\n"
	testBytes := []byte(testInput)

	// Convert byte stream to string
	input := ""
	for i := 0; i < len(testBytes); i++ {
		char := string(testBytes[i])
		input += char
	}

	if len(input) == 0 {
		fmt.Println("Please provide an input")
		return
	}

	switch input[0] {
	case simpleString:
		fmt.Println("Simple string")
	case simpleError:
		fmt.Println("Simple error")
	case integer:
		fmt.Println("Integer")
	case bulkString:
		fmt.Println("Bulk string")
	case array:
		fmt.Println("Array")
	default:
		fmt.Println("Unknown / invalid command")
	}

	fmt.Println("Your command is:", input)
}

func is_CRLF(str string, pointer int) bool {

	// Pointer out of bounds
	if pointer+1 >= len(str) || pointer < 0 {
		return false
	}

	// Check if there is a Carriage Return Line Feed (\r\n)
	if str[pointer] == '\r' && str[pointer+1] == '\n' {
		return true
	}

	return false
}
