// Package set provides a generic interface for mathematical set operations.
// It supports common set operations such as union, intersection, difference, and symmetric difference,
// along with in-place modifications, element iteration, and serialization.
package goset

import (
	"fmt"
	"iter"
)

// Set represents a collection of unique elements of type T.
// It provides mathematical set operations and is compatible with Go's iteration protocol.
//
// The zero value is a nil set, which should be initialized via constructor functions.
type Set[T comparable] interface {
	// Add inserts the element into the set.
	// If the element already exists, it has no effect.
	Add(element T)

	// Remove deletes the element from the set.
	// If the element doesn't exist, it has no effect.
	Remove(any T)

	// Contains reports whether the element exists in the set.
	Contains(any T) bool

	// Union returns a new set containing all elements present in either set.
	Union(other Set[T]) Set[T]

	// Intersection returns a new set containing elements present in both sets.
	Intersection(otherset Set[T]) Set[T]

	// Difference returns a new set containing elements in this set but not in the other.
	Difference(other Set[T]) Set[T]

	// SymmetricDifference returns a new set containing elements present in exactly one set.
	SymmetricDifference(other Set[T]) Set[T]

	// Merge adds all elements from the other set to this set (in-place union).
	Merge(other Set[T])

	// Retain keeps only elements present in both sets (in-place intersection).
	Retain(other Set[T])

	// Subtract removes all elements present in the other set from this set (in-place difference).
	Subtract(other Set[T])

	// Xor replaces this set with elements present in exactly one set (in-place symmetric difference).
	Xor(other Set[T])

	// Equals reports whether two sets contain identical elements.
	Equals(other Set[T]) bool

	// IsSuperset reports whether this set contains all elements of the other set.
	IsSuperset(other Set[T]) bool

	// IsSubset reports whether all elements of this set are present in the other set.
	IsSubset(other Set[T]) bool

	// Elements returns a slice containing all set elements.
	// The order is undefined and may change between realizations.
	Elements() []T

	// All returns an iterator for ranging over elements.
	// Requires GOEXPERIMENT=rangefunc in Go 1.22+.
	All() iter.Seq[T]

	// Len returns the number of elements in the set.
	Len() int

	// String returns a human-readable representation in the format "Set{e1, e2, ...}".
	fmt.Stringer
}
