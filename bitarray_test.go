package bitarray_test

import (
	"testing"

	"github.com/saphal1998/bloomf/bitarray"
)

func TestBitArraySize(t *testing.T) {
	bitArray := bitarray.NewBitArray(100)

	size := bitArray.Size()

	if size != 100 {
		t.Errorf("expected-%v,  got=%v", 100, size)
	}
}

func TestSetBitOutOfBounds(t *testing.T) {
	bitArray := bitarray.NewBitArray(100)

	offset := 1000
	_, err := bitArray.Get(offset)
	if err == nil {
		t.Errorf("expected=err, got=%v", err)
	}

	err = bitArray.Set(offset)
	if err == nil {
		t.Errorf("expected=err, got=%v", err)
	}

	err = bitArray.Unset(offset)
	if err == nil {
		t.Errorf("expected=err, got=%v", err)
	}
}

func TestSetBitWithinBounds(t *testing.T) {
	bitArray := bitarray.NewBitArray(100)

	offset := 50
	initial, err := bitArray.Get(offset)
	if initial || err != nil {
		t.Errorf("expected=%v, got=%v, err=%v", false, initial, err)
	}

	err = bitArray.Set(offset)
	if err != nil {
		t.Errorf("expected=%v, got=%v", nil, err)
	}

	final, err := bitArray.Get(offset)
	if !final {
		t.Errorf("expected=%v, got=%v, err=%v", true, final, err)
	}
}

func TestUnSetBitWithinBounds(t *testing.T) {
	bitArray := bitarray.NewBitArray(100)

	offset := 50
	initial, err := bitArray.Get(offset)
	if initial || err != nil {
		t.Errorf("expected=%v, got=%v, err=%v", false, initial, err)
	}

	err = bitArray.Set(offset)
	if err != nil {
		t.Errorf("expected=%v, got=%v", nil, err)
	}

	intermediate, err := bitArray.Get(offset)
	if !intermediate {
		t.Errorf("expected=%v, got=%v, err=%v", true, intermediate, err)
	}

	err = bitArray.Unset(offset)
	if err != nil {
		t.Errorf("expected=%v, got=%v", nil, err)
	}

	final, err := bitArray.Get(offset)
	if final || err != nil {
		t.Errorf("expected=%v, got=%v, err=%v", false, final, err)
	}
}

func TestToggleWithinBounds(t *testing.T) {
	bitArray := bitarray.NewBitArray(100)

	offset := 50
	initial, err := bitArray.Get(offset)
	if initial || err != nil {
		t.Errorf("expected=%v, got=%v, err=%v", false, initial, err)
	}

	err = bitArray.Set(offset)
	if err != nil {
		t.Errorf("expected=%v, got=%v", nil, err)
	}

	intermediate, err := bitArray.Get(offset)
	if !intermediate {
		t.Errorf("expected=%v, got=%v, err=%v", true, intermediate, err)
	}

	err = bitArray.Toggle(offset)
	if err != nil {
		t.Errorf("expected=%v, got=%v", nil, err)
	}

	final, err := bitArray.Get(offset)
	if final || err != nil {
		t.Errorf("expected=%v, got=%v, err=%v", false, final, err)
	}
}
