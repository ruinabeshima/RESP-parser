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

type Value struct {
	Type   byte   // '+', '-', ':', '$', '*'
	Str    string // Simple string, simple error, bulk string
	Int    int
	Array  []Value
	IsNull bool // Null bulk strings, null array
}

func main() {
	testInput := "$11\r\nhello world\r\n"
	testBytes := []byte(testInput)
}

func parse(data []byte) (Value, error) {
	if len(data) == 0 {
		return Value{}, errors.New("empty payload")
	}

	switch data[0] {
	case simpleString:
		str, err := parseSimpleString(data)
		return Value{Type: simpleString, Str: str}, err

	case simpleError:
		errStr, err := parseSimpleError(data)
		return Value{Type: simpleError, Str: errStr}, err

	case integer:
		num, err := parseInteger(data)
		return Value{Type: integer, Int: num}, err

	case bulkString:
		bstr, err := parseBulkString(data)
		return Value{Type: bulkString, Str: bstr}

	case array:
		fmt.Println("Array")
	default:
		return Value{}, errors.New("unknown / invalid command")
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
Finds \r\n and returns the slice of bytes from "start" up to (but not including) \r\n
For simple strings, simple errors, integers
*/
func readLine(data []byte, start int) ([]byte, error) {
	end := start
	for end < len(data) && !is_CRLF(data, end) {
		end += 1
	}

	// \r\n not included
	if end == len(data) {
		return nil, errors.New("CRLF not included")
	}

	return data[start:end], nil
}

func parseSimpleString(data []byte) (string, error) {

	// Verify prefix
	if data[0] != '+' {
		return "", errors.New("wrong command type")
	}

	// Retrieve command slice
	slice, err := readLine(data, 1)
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
	slice, err := readLine(data, 1)
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
	slice, err := readLine(data, 1)
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

// 　Bool return value is for isNull (null bulk string)
func parseBulkString(data []byte) (string, bool, error) {

	// Verify prefix
	if data[0] != '$' {
		return "", false, errors.New("wrong command type")
	}

	// Retrieve string length and convert to int
	length, err := readLine(data, 1)
	if err != nil {
		return "", false, fmt.Errorf("%w\n", err)
	}
	intLength, err := strconv.Atoi(string(length))
	if err != nil {
		return "", false, fmt.Errorf("%w\n", err)
	}

	// Null bulk strings (-1)
	if intLength == -1 {
		return "", true, nil
	}

	// Calculate where the payload starts and ends
	bulkStart := 1 + len(length) + 2
	bulkEnd := bulkStart + intLength
	if bulkEnd+2 > len(data) {
		return "", false, errors.New("incomplete bulk string payload")
	}

	// Slice the payload directly, and verify CRLF after
	bulkBytes := data[bulkStart:bulkEnd]
	if !is_CRLF(data, bulkEnd) {
		return "", false, errors.New("missing trailing CSLF after bulk string")
	}

	return string(bulkBytes), nil
}
