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

	switch testBytes[0] {
	case simpleString:
		fmt.Println("Simple string")

		if len(testBytes) == 1 {
			fmt.Println("No command")
			return
		}

		end := 1
		for end < len(testBytes) && !is_CRLF(testBytes, end) {
			end += 1
		}
		command := string(testBytes[1:end])
		fmt.Println("Command: ", command)

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

}

func is_CRLF(byteArray []byte, pointer int) bool {

	// Pointer out of bounds
	if pointer+1 >= len(byteArray) || pointer < 0 {
		return false
	}

	// Check if there is a Carriage Return Line Feed (\r\n)
	if byteArray[pointer] == '\r' && byteArray[pointer+1] == '\n' {
		return true
	}

	return false
}
