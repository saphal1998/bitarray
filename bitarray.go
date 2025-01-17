package bitarray

import (
	"bytes"
	"errors"
)

var outOfBounds = errors.New("out of bounds")

type BitArray interface {
	Set(idx int) error
	Unset(idx int) error
	Get(idx int) (bool, error)
	Size() int
	Toggle(idx int) error
	UnsafeRawBuffer() []byte
}

type bitarray struct {
	size   int
	buffer []byte
}

func (b *bitarray) Set(idx int) error {
	buffer_idx := idx >> 3
	if buffer_idx > len(b.buffer) {
		return outOfBounds
	}
	reminder := 7 - (idx & 0b0111)

	correct_blob := b.buffer[buffer_idx]
	correct_bit := byte(0b10000000 >> reminder)

	correct_blob = (correct_blob | correct_bit)
	b.buffer[buffer_idx] = correct_blob
	return nil
}

func (b *bitarray) Unset(idx int) error {
	buffer_idx := idx >> 3
	if buffer_idx > len(b.buffer) {
		return outOfBounds
	}
	reminder := 7 - (idx & 0b0111)

	correct_blob := b.buffer[buffer_idx]
	correct_bit := byte(0b10000000 >> reminder)
	mask := byte(0b11111111) & (^correct_bit)

	correct_blob = (correct_blob & mask)
	b.buffer[buffer_idx] = correct_blob
	return nil
}

func (b *bitarray) Size() int {
	return b.size
}

func (b *bitarray) Toggle(idx int) error {
	status, err := b.Get(idx)
	if err != nil {
		return err
	}

	if status {
		return b.Unset(idx)
	} else {
		return b.Set(idx)
	}

}

func (b *bitarray) Get(idx int) (bool, error) {
	buffer_idx := idx >> 3
	if buffer_idx > len(b.buffer) {
		return false, outOfBounds
	}
	reminder := 7 - (idx & 0b0111)

	correct_blob := b.buffer[buffer_idx]
	correct_bit := byte(0b10000000 >> reminder)

	return (correct_blob & correct_bit) != byte(0), nil
}

func (b *bitarray) UnsafeRawBuffer() []byte {
	return bytes.Clone(b.buffer)
}

func NewBitArray(size int) BitArray {
	reminder := size & 0b0111
	bufferSize := size >> 3
	if reminder != 0 {
		bufferSize += 1
	}
	return &bitarray{
		size:   size,
		buffer: make([]byte, int64(bufferSize)),
	}

}
