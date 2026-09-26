package main

import (
	"flag"
	"fmt"
	"github.com/mschilli/go-murmur"
	"os"
)

func main() {
	filePath := flag.String("file", "", "path to the secrets file (default ~/.murmur)")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options] secret\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	m := murmur.NewMurmur().WithFilePath(*filePath)

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
