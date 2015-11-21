package registrar

import "testing"

func newTestRegistrar() Registrar {
	return NewInmem()
}

func TestReserve(t *testing.T) {
	r := newTestRegistrar()

	obj := "test1"
	if err := r.Reserve("test", obj); err != nil {
		t.Fatal(err)
	}

	if err := r.Reserve("test", obj); err == nil {
		t.Fatal(err)
	}

	obj2 := "test2"
	err := r.Reserve("test", obj2)
	if err == nil {
		t.Fatalf("expected error when reserving an already reserved name to another object")
	}
	if err != ErrNameReserved {
		t.Fatal("expected `ErrNameReserved` error when attempting to reserve an already reserved name")
	}
}

func TestRelease(t *testing.T) {
	r := newTestRegistrar()
	obj := "testing"

	if err := r.Reserve("test", obj); err != nil {
		t.Fatal(err)
	}
	r.Release("test")
	r.Release("test") // Ensure there is no panic here

	if err := r.Reserve("test", obj); err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	r := newTestRegistrar()
	obj := "testing"
	names := []string{"test1", "test2"}
	for _, name := range names {
		if err := r.Reserve(name, obj); err != nil {
			t.Fatal(err)
		}
	}

	r.Reserve("test3", "other")
	r.Delete(obj)

	_, err := r.Get(obj)
	if err == nil {
		t.Fatal("expected error getting names for deleted key")
	}

	if err != ErrNoSuchKey {
		t.Fatalf("expected `ErrNoSuchKey`, got: %v", err)
	}
}

func TestGet(t *testing.T) {
	r := newTestRegistrar()

	if err := r.Reserve("val", "test"); err != nil {
		t.Fatal(err)
	}

	v, err := r.Get("test")
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(v))
	}
	if v[0] != "val" {
		t.Fatalf("expected `val`, got: %s", v)
	}
}
