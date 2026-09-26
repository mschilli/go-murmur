package main

import (
	"flag"
	"fmt"
	"github.com/mschilli/go-murmur"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	filePath := flag.String("file", "", "path to the secrets file (default ~/.murmur)")
	mask := flag.Bool("mask", false, "print the secret key with a masked value")
	list := flag.Bool("list", false, "print all secret keys and values")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options] secret\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "       %s [options] -list\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	m := murmur.NewMurmur().WithFilePath(*filePath)

	if *list {
		if flag.NArg() != 0 {
			flag.Usage()
			os.Exit(1)
		}
		if err := m.Read(); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read secrets file: %v\n", err)
			os.Exit(1)
		}
		keys := make([]string, 0, len(m.Dict))
		for key := range m.Dict {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			value := m.Dict[key]
			if *mask {
				value = strings.Repeat("*", utf8.RuneCountInString(value))
			}
			fmt.Printf("%s: %s\n", key, value)
		}
		return
	}

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
