package murmur

import (
	"testing"
)

func TestLookup(t *testing.T) {
	mur := NewMurmur().WithFilePath("data/murmur.yaml")

	name := "foo"
	p, err := mur.Lookup(name)

	if err != nil {
		t.Log("name", name, "not found")
		t.Fail()
	}
	if p != "bar" {
		t.Log("name", name, "p", p, "mismatch")
		t.Fail()
	}

	name = "nonexist"
	p, err = mur.Lookup(name)

	if err == nil {
		t.Log("name", name, "found")
		t.Fail()
	}
}
