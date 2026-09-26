package main

import (
	"flag"
	"fmt"
	"github.com/mschilli/go-murmur"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	filePath := flag.String("file", "", "path to the secrets file (default ~/.murmur)")
	mask := flag.Bool("mask", false, "print the secret key with a masked value")
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

	if *mask {
		fmt.Printf("%s: %s\n", flag.Arg(0), strings.Repeat("*", utf8.RuneCountInString(secret)))
		return
	}
	fmt.Printf("%s\n", secret)
}
