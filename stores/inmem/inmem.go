// Package inmem stores key/value pairs in memory
package inmem

import (
	"sync"

	"github.com/cpuguy83/registrar/stores"
)

var errNoSuchKey = stores.ErrNoSuchKey

// New creates a new inmem store
func New() stores.Store {
	return &inmem{
		data:   make(map[string][]string),
		valIdx: make(map[string]string),
	}
}

// inmem stores k/v's in memory
type inmem struct {
	data   map[string][]string
	valIdx map[string]string
	mu     sync.Mutex
}

// Get returns the list of values assigned to a key
func (m *inmem) Get(key string) ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, exists := m.data[key]
	if !exists {
		return nil, errNoSuchKey
	}
	return v, nil
}

// Exists returns a bool indicating if the value exists for any stored key
func (m *inmem) Exists(value string) bool {
	m.mu.Lock()
	_, exists := m.valIdx[value]
	m.mu.Unlock()
	return exists
}

// Set sets the value to the given key
func (m *inmem) Set(key, value string) error {
	m.mu.Lock()
	m.data[key] = append(m.data[key], value)
	m.valIdx[value] = key
	m.mu.Unlock()
	return nil
}

// Delete deletes all entries for a given key
func (m *inmem) Delete(key string) error {
	m.mu.Lock()
	vals := m.data[key]
	for _, v := range vals {
		delete(m.valIdx, v)
	}
	delete(m.data, key)
	m.mu.Unlock()
	return nil
}

// DeleteValue deletes the stored value
func (m *inmem) DeleteValue(value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	key, exists := m.valIdx[value]
	if !exists {
		return nil
	}

	for i, v := range m.data[value] {
		if v != value {
			continue
		}
		m.data[key] = append(m.data[key][:i], m.data[key][i+1:]...)
		break
	}

	delete(m.valIdx, value)
	if len(m.data[key]) == 0 {
		delete(m.data, key)
	}
	return nil
}
