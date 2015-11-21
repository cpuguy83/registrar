// Package registrar provides name registration/reservation.
// It reserves a name to a given key. Keys can have multiple names, but names are
// unique across all keys.
package registrar

import "errors"

var (
	// ErrNameReserved is an error which is returned when a name is requested to be reserved that already is reserved
	ErrNameReserved = errors.New("name is reserved")
	// ErrNoSuchKey is returned when trying to find the names for a key which is not known
	ErrNoSuchKey = errors.New("provided key does not exist")
)

// Registrar stores a list of keys and their registered names as well as indexes names and the key that they are registred to
// Names must be unique.
type Registrar interface {
	// Reserve registers a key to a name
	// Reserve is idempotent
	// Attempting to reserve a key to a name that already exists reults in an `ErrNameReserved`
	Reserve(name, key string) error
	// Release releases the reserved name
	// Once released, a name can be be reserved again
	Release(name string) error
	// Delete removes all reservations for the passed in key.
	// All names reserved to this key are released.
	Delete(key string) error
	// Get lists all the reserved names for the given key
	Get(key string) ([]string, error)
	// List lists all reserved names and the keys they are associated with
	List() (map[string][]string, error)
}
