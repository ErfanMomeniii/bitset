package bitset

import (
	"errors"
	"reflect"
)

var (
	ErrorShouldHaveSameLength  = errors.New("bitsets should have same length")
	ErrorCannotParseFromString = errors.New("cannot parse bitset from input string")
	ErrorCannotCreateBitSet    = errors.New("cannot create bitset from input parameter")
)

type BitOperation interface {
	Xor(bitset1 Bitset, bitset2 Bitset) (Bitset, error)
	And(bitset1 Bitset, bitset2 Bitset) (Bitset, error)
	Or(bitset1 Bitset, bitset2 Bitset) (Bitset, error)
}

type BitSetOperation interface {
	Not()
	Count() int
	Set(index int)
	Reset(index int)
	Flip(index int)
	None() bool
}

type Bitset []bool

func New(t interface{}) (*Bitset, error) {
	v := reflect.ValueOf(t)

	if v.Kind() == reflect.Int {
		b := make(Bitset, v.Int())

		return &b, nil
	}

	if v.Kind() == reflect.String {
		b, err := ParseFromString(v.String())

		if err != nil {
			return nil, err
		}

		return &b, nil
	}

	return nil, ErrorCannotCreateBitSet
}

func (b *Bitset) Not() {
	for _, bit := range interface{}(b).([]bool) {
		bit = !bit
	}
	return
}

func (b *Bitset) Count() int {
	var result int

	for _, bit := range interface{}(b).([]bool) {
		if bit {
			result++
		}
	}

	return result
}

func (b *Bitset) Set(index int) {
	interface{}(b).([]bool)[index] = true
}

func (b *Bitset) Reset(index int) {
	interface{}(b).([]bool)[index] = false
}

func (b *Bitset) Flip(index int) {
	interface{}(b).([]bool)[index] = !interface{}(b).([]bool)[index]
}

func (b *Bitset) None() bool {
	return b.Count() == 0
}

func ParseFromString(s string) (Bitset, error) {
	result := make([]bool, len(s))

	for i, c := range s {
		switch c {
		case '1':
			result[i] = true
		case '0':
			result[i] = false
		default:
			return nil, ErrorCannotParseFromString
		}
	}

	return result, nil
}

func Xor(bitset1 Bitset, bitset2 Bitset) (Bitset, error) {
	l1 := len(bitset1)
	l2 := len(bitset2)
	if l1 != l2 {
		return nil, ErrorShouldHaveSameLength
	}

	result, _ := New(l1)

	for i := 0; i < l1; i++ {
		a1 := interface{}(bitset1[i]).(int)
		a2 := interface{}(bitset2[i]).(int)
		ans := a1 ^ a2
		interface{}(result).([]bool)[i] = interface{}(ans).(bool)
	}

	return *result, nil
}

func And(bitset1 Bitset, bitset2 Bitset) (Bitset, error) {
	l1 := len(bitset1)
	l2 := len(bitset2)
	if l1 != l2 {
		return nil, ErrorShouldHaveSameLength
	}

	result, _ := New(l1)

	for i := 0; i < l1; i++ {
		a1 := interface{}(bitset1[i]).(int)
		a2 := interface{}(bitset2[i]).(int)
		ans := a1 & a2
		interface{}(result).([]bool)[i] = interface{}(ans).(bool)
	}

	return *result, nil
}

func Or(bitset1 Bitset, bitset2 Bitset) (Bitset, error) {
	l1 := len(bitset1)
	l2 := len(bitset2)
	if l1 != l2 {
		return nil, ErrorShouldHaveSameLength
	}

	result, _ := New(l1)

	for i := 0; i < l1; i++ {
		a1 := interface{}(bitset1[i]).(int)
		a2 := interface{}(bitset2[i]).(int)
		ans := a1 | a2
		interface{}(result).([]bool)[i] = interface{}(ans).(bool)
	}

	return *result, nil
}
