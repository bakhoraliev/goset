package goset

import (
	"iter"
	"sync"
)

// SyncSet is a thread-safe wrapper around any Set implementation.
// It uses RWMutex to separate read and write operations for better concurrency.
//
// The zero value is a nil set, which should be initialized via NewSyncSet.
type SyncSet[T comparable] struct {
	Set[T]
	mu sync.RWMutex
}

// NewSyncSet creates a new SyncSet wrapping the provided Set implementation.
func NewSyncSet[T comparable](set Set[T]) *SyncSet[T] {
	return &SyncSet[T]{
		Set: set,
	}
}

// Add inserts the element into the set. Write-locked.
func (set *SyncSet[T]) Add(element T) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.Set.Add(element)
}

// Remove deletes the element from the set. Write-locked.
func (set *SyncSet[T]) Remove(element T) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.Set.Remove(element)
}

// Contains reports whether the element exists in the set. Read-locked.
func (set *SyncSet[T]) Contains(element T) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Contains(element)
}

// Union returns a new set containing all elements present in either set.
// The returned set is NOT thread-safe. Read-locked.
func (set *SyncSet[T]) Union(other Set[T]) Set[T] {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Union(other)
}

// Intersection returns a new set containing elements present in both sets.
// The returned set is NOT thread-safe. Read-locked.
func (set *SyncSet[T]) Intersection(other Set[T]) Set[T] {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Intersection(other)
}

// Difference returns a new set containing elements in this set but not in the other.
// The returned set is NOT thread-safe. Read-locked.
func (set *SyncSet[T]) Difference(other Set[T]) Set[T] {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Difference(other)
}

// SymmetricDifference returns a new set containing elements present in exactly one set.
// The returned set is NOT thread-safe. Read-locked.
func (set *SyncSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.SymmetricDifference(other)
}

// Merge adds all elements from the other set to this set (in-place union). Write-locked.
func (set *SyncSet[T]) Merge(other Set[T]) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.Set.Merge(other)
}

// Retain keeps only elements present in both sets (in-place intersection). Write-locked.
func (set *SyncSet[T]) Retain(other Set[T]) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.Set.Retain(other)
}

// Subtract removes all elements present in the other set from this set (in-place difference). Write-locked.
func (set *SyncSet[T]) Subtract(other Set[T]) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.Set.Subtract(other)
}

// Xor replaces this set with elements present in exactly one set (in-place symmetric difference). Write-locked.
func (set *SyncSet[T]) Xor(other Set[T]) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.Set.Xor(other)
}

// Equals reports whether two sets contain identical elements. Read-locked.
func (set *SyncSet[T]) Equals(other Set[T]) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Equals(other)
}

// IsSuperset reports whether this set contains all elements of the other set. Read-locked.
func (set *SyncSet[T]) IsSuperset(other Set[T]) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.IsSuperset(other)
}

// IsSubset reports whether all elements of this set are present in the other set. Read-locked.
func (set *SyncSet[T]) IsSubset(other Set[T]) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.IsSubset(other)
}

// Elements returns a slice containing all set elements.
// The order is undefined and may change between realizations. Read-locked.
func (set *SyncSet[T]) Elements() []T {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Elements()
}

// Clone returns a copy of the set.
// The clone is NOT thread-safe. Read-locked.
func (set *SyncSet[T]) Clone() Set[T] {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Clone()
}

// All returns an iterator for ranging over elements.
// IMPORTANT: Iteration is NOT thread-safe. The iterator must be consumed
// while the lock is held, or the underlying set must not be modified during iteration. Read-locked.
func (set *SyncSet[T]) All() iter.Seq[T] {
	set.mu.RLock()
	return func(yield func(T) bool) {
		defer set.mu.RUnlock()
		for item := range set.Set.All() {
			if !yield(item) {
				return
			}
		}
	}
}

// AllSafe returns an iterator over a snapshot of the set elements.
// The snapshot is taken under read lock, so iteration is thread-safe
// and will not see concurrent modifications.
//
// Time complexity: O(n) where n is the size of the set (for Elements() call). Read-locked.
func (set *SyncSet[T]) AllSafe() iter.Seq[T] {
	set.mu.RLock()
	snapshot := set.Set.Elements()
	set.mu.RUnlock()
	return func(yield func(T) bool) {
		for _, item := range snapshot {
			if !yield(item) {
				return
			}
		}
	}
}

// Len returns the number of elements in the set. Read-locked.
func (set *SyncSet[T]) Len() int {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.Len()
}

// String returns a human-readable representation in the format "Set{e1, e2, ...}". Read-locked.
func (set *SyncSet[T]) String() string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.Set.String()
}
