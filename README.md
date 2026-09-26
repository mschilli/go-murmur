# Murmur

Go library for simple access to a YAML formatted password file.

## How to use it

```
$ cat ~/.murmur
fooapp: topsecret
barapp: hunter3
```

By default, Murmur reads secrets from `~/.murmur`. The file contents are YAML,
but the default filename does not use a `.yaml` extension. Use `WithFilePath()`
to read a different file path.

For the command-line tool, use `-file` before the secret name:

```sh
murmur -file /path/to/other.murmur fooapp
```

```
$ cat mtest.go
```

<!--(Config::Patch-example1-replace)-->
```go
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
```
<!--(Config::Patch::replace)-->
<!-- RVhBTVBMRTEK-->
<!--(Config::Patch::replace)-->
<!--(Config::Patch-example1-replace)-->

## API docs

<!--(Config::Patch-apiusage-replace)-->
```
package murmur // import "github.com/mschilli/go-murmur"

type Murmur struct {
	FilePath string

	Dict map[string]string
	// Has unexported fields.
}
    Read secrets from a .murmur YAML file

func NewMurmur() *Murmur
    Create a new instance

func (m *Murmur) Lookup(name string) (string, error)
    Look up a .murmur key by name and return its value

func (m *Murmur) Read() error
    Read the .murmur file into the internal cache

func (m *Murmur) WithFilePath(path string) *Murmur
    Set the .murmur file path manually

```
<!--(Config::Patch::replace)-->
<!-- QVBJVVNBR0UK-->
<!--(Config::Patch::replace)-->
<!--(Config::Patch-apiusage-replace)-->

## Author

Mike Schilli, m@perlmeister.com 2024

## License

Released under the [Apache 2.0](LICENSE)
