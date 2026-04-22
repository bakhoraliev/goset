package goset

import (
	"fmt"
	"iter"
	"maps"
	"strings"
)

// HashSet is a map-based implementation of a Set. It uses a map[T]struct{} for storage,
// providing O(1) time complexity for basic operations like Add, Remove, and Contains.
// The zero value is NOT usable - use NewHashSet() to create instances.
type HashSet[T comparable] map[T]struct{}

// NewHashSet creates a new empty HashSet. Always use this constructor to initialize the set.
func NewHashSet[T comparable]() HashSet[T] {
	return make(HashSet[T])
}

// Add inserts an element into the set.
// If the element already exists, it's a no-op.
//
// Time complexity: O(1).
func (set HashSet[T]) Add(element T) {
	set[element] = struct{}{}
}

// Remove deletes an element from the set.
// If the element doesn't exist, it's a no-op.
//
// Time complexity: O(1).
func (set HashSet[T]) Remove(element T) {
	delete(set, element)
}

// Contains returns true if the element exists in the set.
//
// Time complexity: O(1)
func (set HashSet[T]) Contains(element T) bool {
	_, ok := set[element]
	return ok
}

// Union returns a new set containing all elements present in either set.
//
// Time complexity: O(n + m) where n and m are the sizes of the sets.
func (set HashSet[T]) Union(other Set[T]) Set[T] {
	newset := maps.Clone(set)
	for element := range other.All() {
		newset.Add(element)
	}
	return newset
}

// Intersection returns a new set containing elements present in both sets.
//
// Time complexity: O(n) where n is size of the _other_ set.
func (set HashSet[T]) Intersection(other Set[T]) Set[T] {
	newset := NewHashSet[T]()
	for element := range other.All() {
		if set.Contains(element) {
			newset.Add(element)
		}
	}
	return newset
}

// Difference returns a new set containing elements in this set but not in the other.
//
// Time complexity: O(n * c) where n is size of the _current_ set and c is time complexity of the other set's Contains() method.
//
// For two HashSet implementations, this operates in O(n) average time, as Contains() is O(1).
// If the other set has O(m) Contains() complexity (where m = its size), total complexity becomes O(n*m).
func (set HashSet[T]) Difference(other Set[T]) Set[T] {
	newset := NewHashSet[T]()
	for element := range set.All() {
		if !other.Contains(element) {
			newset.Add(element)
		}
	}
	return newset
}

// SymmetricDifference returns a new set containing elements present in exactly one set.
//
// Time complexity: O(n + m) where n and m are the sizes of the sets.
func (set HashSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	newset := maps.Clone(set)
	for element := range other.All() {
		if newset.Contains(element) {
			newset.Remove(element)
		} else {
			newset.Add(element)
		}
	}
	return newset
}

// Merge adds all elements from the other set to this set (in-place union).
//
// Time complexity: O(n) where n is size of the _other_ set.
func (set HashSet[T]) Merge(other Set[T]) {
	for element := range other.All() {
		set.Add(element)
	}
}

// Retain keeps only elements present in both sets (in-place intersection).
//
// Time complexity: O(n * c) where n is size of the _current_ set and c is time complexity of the other set's Contains() method.
//
// For two HashSet implementations, this operates in O(n) average time, as Contains() is O(1).
// If the other set has O(m) Contains() complexity (where m = its size), total complexity becomes O(n*m).
func (set HashSet[T]) Retain(other Set[T]) {
	for item := range set {
		if !other.Contains(item) {
			set.Remove(item)
		}
	}
}

// Subtract removes all elements present in the other set from this set (in-place difference).
//
// Time complexity: O(n * c) where n is size of the _current_ set and c is time complexity of the other set's Contains() method.
//
// For two HashSet implementations, this operates in O(n) average time, as Contains() is O(1).
// If the other set has O(m) Contains() complexity (where m = its size), total complexity becomes O(n*m).
func (set HashSet[T]) Subtract(other Set[T]) {
	for item := range set {
		if other.Contains(item) {
			set.Remove(item)
		}
	}
}

// Xor replaces this set with elements present in exactly one set (in-place symmetric difference).
//
// Time complexity: O(n) where n is size of the _other_ set.
func (set HashSet[T]) Xor(other Set[T]) {
	for element := range other.All() {
		if set.Contains(element) {
			set.Remove(element)
		} else {
			set.Add(element)
		}
	}
}

// Equals reports whether two sets contain identical elements.
//
// Time complexity: O(l + (n * c)) where l is time complexity of the _other_ set's All() method and n is size of the _current_ set and c is time complexity of the _other_ set's Contains() method.
//
// For two HashSet implementations, this operates in O(n) average time, as Len() is O(1) and Contains() is O(1).
func (set HashSet[T]) Equals(other Set[T]) bool {
	if set.Len() != other.Len() {
		return false
	}
	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSuperset reports whether this set contains all elements of the other set.
//
// Time complexity: O(l + (n * c)) where l is time complexity of the _other_ set's Len() method and n is size of the _current_ set and c is time complexity of the _other_ set's Contains() method.
//
// For two HashSet implementations, this operates in O(n) average time, as Len() is O(1) and Contains() is O(1).
func (set HashSet[T]) IsSuperset(other Set[T]) bool {
	if set.Len() < other.Len() {
		return false
	}
	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSubset reports whether all elements of this set are present in the other set.
//
// Time complexity: O(l + (n * c)) where l is time complexity of the _other_ set's Len() method and n is size of the _current_ set and c is time complexity of the _other_ set's Contains() method.
//
// For two HashSet implementations, this operates in O(n) average time, as Len() is O(1) and Contains() is O(1).
func (set HashSet[T]) IsSubset(other Set[T]) bool {
	if set.Len() > other.Len() {
		return false
	}
	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// Elements returns a slice containing all set elements.
// The order of elements is undefined and may vary between calls.
func (set HashSet[T]) Elements() []T {
	elements := make([]T, 0, len(set))
	for item := range set {
		elements = append(elements, item)
	}
	return elements
}

// All returns an iterator for ranging over elements.
// Requires GOEXPERIMENT=rangefunc in Go 1.22+.
func (set HashSet[T]) All() iter.Seq[T] {
	return maps.Keys(set)
}

// Len returns the number of elements in the set.
func (set HashSet[T]) Len() int {
	return len(set)
}

// String returns a human-readable representation in the format "Set{e1, e2, ...}".
func (set HashSet[T]) String() string {
	elements := make([]string, 0, len(set))
	for item := range set {
		elements = append(elements, fmt.Sprintf("%v", item))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(elements, ", "))
}
