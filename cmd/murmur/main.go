package main

import (
	"flag"
	"fmt"
	"github.com/mschilli/go-murmur"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("usage: %s secret\n", os.Args[0])
	}
	flag.Parse()

	m := murmur.NewMurmur()

	if flag.NArg() != 1 {
		flag.Usage()
		return
	}

	secret, err := m.Lookup(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot find secret '%s'\n", flag.Arg(0))
		os.Exit(1)
	}

	fmt.Printf("%s\n", secret)
}
