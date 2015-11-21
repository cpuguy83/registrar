package registrar

import "sync"

// NewInmem creates a new inmem Registrar
func NewInmem() Registrar {
	return &inmem{
		data:   make(map[string][]string),
		valIdx: make(map[string]string),
	}
}

// inmem implements Registrar and stores everything in memory
// inmem is safe for concurrent access
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
		return nil, ErrNoSuchKey
	}
	return v, nil
}

// Reserve sets the value to the given key
func (m *inmem) Reserve(name, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.valIdx[name]; exists {
		return ErrNameReserved
	}
	m.data[key] = append(m.data[key], name)
	m.valIdx[name] = key
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

// Release deletes the stored value
func (m *inmem) Release(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	key, exists := m.valIdx[name]
	if !exists {
		return nil
	}

	for i, v := range m.data[name] {
		if v != name {
			continue
		}
		m.data[key] = append(m.data[key][:i], m.data[key][i+1:]...)
		break
	}

	delete(m.valIdx, name)
	if len(m.data[key]) == 0 {
		delete(m.data, key)
	}
	return nil
}

// List lists all stored items
func (m *inmem) List() (map[string][]string, error) {
	m.mu.Lock()

	// copy items into new map
	ls := make(map[string][]string)
	for k, v := range m.data {
		ls[k] = v
	}

	m.mu.Unlock()
	return ls, nil
}
