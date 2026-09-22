package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Data types correspond to symbol of first byte
const (
	simpleString = '+'
	simpleError  = '-'
	integer      = ':'
	bulkString   = '$'
	array        = '*'
)

func main() {
	testInput := ":1000\r\n"
	testBytes := []byte(testInput)

	switch testBytes[0] {
	case simpleString:
		fmt.Println("Simple string")
		command, err := parseSimpleString(testBytes)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Command:", command)

	case simpleError:
		fmt.Println("Simple error")
		message, err := parseSimpleError(testBytes)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Error Message:", message)

	case integer:
		fmt.Println("Integer")

		num, err := parseInteger(testBytes)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Integer:", num)

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

/*
Finds \r\n and returns the slice of bytes up to (but not including) \r\n
For simple strings, simple errors, integers
*/
func readLine(data []byte) ([]byte, error) {
	end := 1
	for end < len(data) && !is_CRLF(data, end) {
		end += 1
	}

	// \r\n not included
	if end == len(data) {
		return nil, errors.New("CRLF not included")
	}

	return data[1:end], nil
}

func parseSimpleString(data []byte) (string, error) {

	// Verify prefix
	if data[0] != '+' {
		return "", errors.New("wrong command type")
	}

	// Retrieve command slice
	slice, err := readLine(data)
	if err != nil {
		return "", fmt.Errorf("%w\n", err)
	}

	command := string(slice)
	return command, nil
}

func parseSimpleError(data []byte) (string, error) {

	// Verify prefix
	if data[0] != '-' {
		return "", errors.New("wrong command type")
	}

	// Retrieve command slice
	slice, err := readLine(data)
	if err != nil {
		return "", fmt.Errorf("%w\n", err)
	}

	command := string(slice)
	return command, nil
}

func parseInteger(data []byte) (int, error) {

	// Verify prefix
	if data[0] != ':' {
		return 0, errors.New("wrong command type")
	}

	// Retrieve command slice
	slice, err := readLine(data)
	if err != nil {
		return 0, fmt.Errorf("%w\n", err)
	}

	// Convert bytes to string, then parse to int
	num, err := strconv.Atoi(string(slice))
	if err != nil {
		return 0, fmt.Errorf("%w\n", err)
	}

	return num, nil
}
