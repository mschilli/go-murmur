# Murmur

Go library for simple access to a YAML formatted password file.

## How to use it

```
$ cat ~/.murmur.yaml
fooapp: topsecret
barapp: hunter3

$ cat mtest.go
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

$ ./mtest
val=hunter3

## Author

Mike Schilli, m@perlmeister.com 2024

## License

Released under the [Apache 2.0](LICENSE)
