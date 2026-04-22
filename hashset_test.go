package goset

import (
	"testing"
	"testing/quick"
)

func TestEmptySetProperties(t *testing.T) {
	emptyset := NewHashSet[int]()

	// // Empty set should have 0 length
	if emptyset.Len() != 0 {
		t.Errorf("Empty set should have length 0, got %d", emptyset.Len())
	}

	// Empty set should not contain any elements
	property := func(x int) bool {
		return !emptyset.Contains(x)
	}
	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}
