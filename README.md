Registrar
=========

Registrar is a library for registering a name to some string identifier (key).
A key can have multiple names attached to it, but a name can only be registered
once.

A in-memory implementation is provided with this package.

## Usage

```go
package main

import "github.com/cpuguy83/registrar/stores/inmem"

func main() {
	reg := NewInmem()
	reg.Reserve("some_name", "someID")
	reg.Release("some_name")

	reservations, _ := reg.Get("someID")
}
```
