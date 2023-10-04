# ByteString

A Go library for converting strings to byte slices and vice versa without memory allocation.

## Installation

To install the `bytestring` library, use the `go get` command:

```shell
go get github.com/AH-dark/bytestring
```

## Usage

The `bytestring` library provides functions to efficiently convert strings to byte slices and vice versa without memory allocation.

### Convert String to Bytes
```go
import "github.com/AH-dark/bytestring"

s := "Hello, World!"
b := bytestring.StringToBytes(s)
```

### Convert Bytes to String
```go
import "github.com/AH-dark/bytestring"

b := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
s := bytestring.BytesToString(b)
```

## Running Tests

To run tests, use the `go test` command in the project directory:

```bash
go test -v
```

Also, you can run benchmarks:

```bash
go test -bench=. -benchmem
```

## License

This project is licensed under the GNU General Public License v3.0 - see the `LICENSE` file for details.
