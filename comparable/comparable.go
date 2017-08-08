package comparable

import "fmt"

// Comparable mimicing Java interface
type Comparable interface {
	CompareTo(n Integer32) int8
}

// Integer32 my special integer type
type Integer32 int32

// T a try to mimic a generic type
type T struct {
	Value Integer32
}

// CompareTo mimic
func (t T) CompareTo(n Integer32) int8 {
	thisVal := t.Value
	anotherVal := n
	fmt.Printf("thisVal = %d; anotherVal = %d\n", thisVal, anotherVal)
	if thisVal < anotherVal {
		return -1
	} else if thisVal > anotherVal {
		return 1
	} else {
		return 0
	}
}
