# RESP2 Parser

A lightweight, zero-dependency Go implementation of a **RESP2 (Redis Serialization Protocol)** parser. Built from scratch, and operates directly on raw byte streams using safe slice operations and recursive decoding.

## Features

* **Binary-Safe Decoding:** Reads Bulk Strings by byte length rather than scanning line-by-line, preserving raw binary content like images or compressed payloads.
* **Zero Allocations for Intermediate Parsing:** Employs byte slicing (`[]byte`) to locate bounds before converting into Go primitives.
* **Recursive Array Handling:** Seamlessly parses nested data structures using byte consumption tracking across elements.
* **Full RESP2 Specification Support:**
    - Simple Strings (`+`)
    - Simple Errors (`-`)
    - Integers (`:`)
    - Bulk Strings (`$`), including empty and `nil` (`$-1\r\n`) representations
    - Arrays (`*`), including recursive/nested arrays and `nil` (`*-1\r\n`) arrays

## Architecture 

### The `Value` Type

The `Value` struct serves as a universal container capable of holding any decoded RESP message:

```go
type Value struct {
	Type   byte    // Prefix symbol: '+', '-', ':', '$', or '*'
	Str    string  // Simple Strings, Simple Errors, and Bulk Strings
	Int    int     // Integers
	Array  []Value // Nested elements for Array payloads
	IsNull bool    // Indicates Redis nil values ($-1 or *-1)
}

```

### Consumption Tracking

Every parser helper returns three parameters: `(data, consumedBytes, error)`.

When parsing an **Array**, the parser iterates through child items, calling `parse(data[offset:])` recursively and incrementing `offset += consumed` until the expected element count is satisfied. This enables handling deeply nested structures without manual buffer management.


## Getting Started

### Prerequisites

* **Go:** 1.18 or higher

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/resp-parser-go.git
cd resp-parser-go
```


