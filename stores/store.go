// Package stores defines the interface used for name reservation backends for the registrator package
package stores

import "errors"

// ErrNoSuchKey is returned when trying to find the names for a key which is not known
var ErrNoSuchKey = errors.New("provided key does not exist")

// Store is used to manage name reservations from the registrar package
type Store interface {
	// Get returns the list of values assigned to a key
	Get(key string) ([]string, error)
	// Exists returns a bool indicating if the value exists for any stored key
	Exists(value string) bool
	// Set sets the value to the given key
	Set(key, value string) error
	// Delete deletes all entries for a given key
	Delete(key string) error
	// DeleteValue deletes the stored value
	DeleteValue(value string) error
}
