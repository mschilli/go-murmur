package main

import (
	"fmt"
	"github.com/mschilli/go-murmur"
	"os"
	"text/template"
)

func main() {
	m := murmur.NewMurmur().WithFilePath("data/murmur")

	// Retrieve a single secret
	val, err := m.Lookup("secret1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("val: %s\n", val)

	// Replace secrets in a template
	t, err := template.ParseFiles("data/test.tmpl")
	err = t.Execute(os.Stdout, m.Dict)
	if err != nil {
		panic(err)
	}
}
