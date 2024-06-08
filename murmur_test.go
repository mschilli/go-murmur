package murmur

import (
    "testing"
)

func TestLookup(t *testing.T) {
    StoreLocationFinder = func() (string, error) { return "data/murmur.yaml", nil }

    name := "foo"
    p, err := Lookup(name)

    if err != nil {
	t.Log("name", name, "not found")
	t.Fail()
    }
    if p != "bar" {
	t.Log("name", name, "p", p, "mismatch")
	t.Fail()
    }

    name = "nonexist"
    p, err = Lookup(name)

    if err == nil {
	t.Log("name", name, "found")
	t.Fail()
    }
}
