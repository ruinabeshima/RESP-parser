# RESP2 Parser

A small, dependency-free Go package that reads **RESP2** (the wire format Redis uses) from raw bytes and turns it into Go values.

Used as the parser in [ruinabeshima/mini-redis](https://github.com/ruinabeshima/mini-redis).

## How it works

`Parse` looks at the first byte to decide what the message is, hands it to the matching helper, and returns the value plus how many bytes it used.

```
 bytes off the connection
            │
            ▼
      Parse(data)  ── looks at data[0]
            │
   ┌────────┼──────────────────────────────┐
   │ '+'  simple string                    │
   │ '-'  error                            │
   │ ':'  integer                          │──► (Value, bytesRead, error)
   │ '$'  bulk string  (read by length)    │
   │ '*'  array ──┐                        │
   └──────────────┼────────────────────────┘
                  │  calls Parse on each element,
                  └─ adding up bytesRead as it goes
                        (nested arrays just work)
```

Every helper returns `bytesRead` so the caller knows where the next message starts in the buffer. If the bytes end mid-message, you get `ErrIncomplete` — read more from the socket and try again.

## The `Value` type

One struct holds any RESP2 message:

```go
type Value struct {
	Type   byte    // '+', '-', ':', '$', or '*'
	Str    string  // simple strings, errors, bulk strings
	Int    int     // integers
	Array  []Value // elements, for arrays
	IsNull bool    // $-1 or *-1
}
```

## Supported types

Simple strings (`+`), errors (`-`), integers (`:`), bulk strings (`$`, including empty and null), and arrays (`*`, including nested and null).

Bulk strings are read by their declared byte length rather than scanned for a newline, so binary payloads pass through untouched.

## Usage

```bash
go get github.com/ruinabeshima/RESP-parser
```

```go
import "github.com/ruinabeshima/RESP-parser/resp"

val, n, err := resp.Parse(buf)
```

Requires Go 1.18+.
