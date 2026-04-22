package goset

import (
	"encoding/json"
)

// MarshalSet is a wrapper that adds JSON marshaling to any Set implementation.
// It wraps an existing Set and implements json.Marshaler and json.Unmarshaler.
//
// The zero value is not usable - use NewMarshalSet to create instances.
type MarshalSet[T comparable] struct {
	Set[T]
	json.Marshaler
	json.Unmarshaler
}

// NewMarshalSet creates a new MarshalSet wrapping the provided Set.
func NewMarshalSet[T comparable](set Set[T]) *MarshalSet[T] {
	return &MarshalSet[T]{Set: set}
}

// MarshalJSON implements json.Marshaler.
// Returns JSON array of set elements.
func (s *MarshalSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Set.Elements())
}

// UnmarshalJSON implements json.Unmarshaler.
// Adds all elements from JSON array to the set. Does NOT clear existing elements.
// To clear before unmarshaling, create a new empty set first.
func (s *MarshalSet[T]) UnmarshalJSON(data []byte) error {
	var elements []T
	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}
	for _, e := range elements {
		s.Set.Add(e)
	}
	return nil
}
