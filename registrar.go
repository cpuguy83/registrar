// Package registrar provides name registration/reservation. It reserves a name to a given key.
package registrar

import (
	"errors"
	"sync"

	"github.com/cpuguy83/registrar/stores"
)

// ErrNameReserved is an error which is returned when a name is requested to be reserved that already is reserved
var (
	ErrNameReserved = errors.New("name is reserved")
	ErrNoSuchKey    = stores.ErrNoSuchKey
)

// Registrar stores indexes a list of keys and their registered names as well as indexes names and the key that they are registred to
// Names must be unique.
// Registrar is safe for concurrent access.
type Registrar struct {
	store stores.Store
	mu    sync.Mutex
}

// NewRegistrar creates a new Registrar with the an empty index
func NewRegistrar(s stores.Store) *Registrar {
	return &Registrar{store: s}
}

// Reserve registers a key to a name
// Reserve is idempotent
// Attempting to reserve a key to a name that already exists reults in an `ErrNameReserved`
// A name reservation is globally unique
func (r *Registrar) Reserve(name, key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.store.Exists(name) {
		return ErrNameReserved
	}
	return r.store.Set(key, name)
}

// Release releases the reserved name
// Once released, a name can be be reserved again
func (r *Registrar) Release(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.store.DeleteValue(name)
}

// Delete removes all reservations for the passed in key.
// All names reserved to this key are released.
func (r *Registrar) Delete(key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.store.Delete(key)
}

// Get lists all the reserved names for the given key
func (r *Registrar) Get(key string) ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.store.Get(key)
}
