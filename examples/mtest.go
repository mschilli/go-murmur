package main

import (
    "fmt"
    "github.com/mschilli/go-murmur"
)

func main() {
    m := murmur.NewMurmur()
    val, err := m.Lookup("barapp")

    if err != nil {
	panic(err)
    }

    fmt.Printf("val: %s\n", val)
}
