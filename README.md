# bitarray

The `bitarray` package provides a compact and efficient way to represent, manipulate, and store bit arrays in Go.

## Installation

```
go get -v github.com/saphal1998/bitarray
```

## Usage

The `bitarray` package defines the `BitArray` interface and a concrete implementation `bitarray`.

### BitArray Interface

The `BitArray` interface specifies the methods for interacting with a bit array.

* `Set(idx int) error`: Sets the bit at the specified index to 1.
* `Unset(idx int) error`: Sets the bit at the specified index to 0.
* `Get(idx int) (bool, error)`: Returns the value of the bit at the specified index and an error if the index is out of bounds.
* `Size() int`: Returns the size of the bit array in bits.
* `Toggle(idx int) error`: Flips the bit at the specified index.

### bitarray struct

The `bitarray` struct implements the `BitArray` interface. It uses a slice of bytes to store the bits efficiently, packing 8 bits into each byte.

* `NewBitArray(size int) BitArray`: Creates a new `bitarray` of the specified size in bits.

## Example

```go
package main

import (
	"fmt"

	"github.com/saphal1998/bitarray"
)

func main() {
	// Create a new bit array of size 10
	ba := bitarray.NewBitArray(10)

	// Set the bit at index 5
	err := ba.Set(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the value of the bit at index 5
	val, err := ba.Get(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bit at index 5:", val) // Output: Bit at index 5: true

	// Toggle the bit at index 3
	err = ba.Toggle(3)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the size of the bit array
	fmt.Println("Size of the bit array:", ba.Size()) // Output: Size of the bit array: 10
}
```

## License

The `bitarray` package is licensed under the MIT License. See the LICENSE file for details.

