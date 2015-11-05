Registrar
=========

Registrar is a library for registering a name to some string identifier (key).
A key can have multiple names attached to it, but a name can only be registered
once.

## Usage

Registrar uses a `Store` interface to be able to swap out different storage
backends. Today there is only one storage backend provided, `inmem`, which stores
key/values in a map in memory.

```go
package main

import "github.com/cpuguy83/registrar/stores/inmem"

func main() {
	reg := NewRegistrar(inmem.New())
	reg.Reserve("some_name", "someID")
	reg.Release("some_name")

	reservations, _ := reg.Get("someID")
}
```

## Known Issues

Locking on the `Registrar` type leaves some room for imrpovement. It's fine for
fast access like `inmem`'s map, but with a slower backend, the current locking
could slow things down.
