package ref_test

import (
    "fmt"
    ref "github.com/oas3/json-reference"
)

func ExampleNew_fragment() {
    r, _ := ref.New("#/ok")
    fmt.Println(r.URL.String())
    fmt.Println(r.Pointer.String())

    // Output:
    // #/ok
    // /ok
}

func ExampleNew_full() {
    r, _ := ref.New("http://example.com/example.json#/foo/bar")
    fmt.Println(r.URL.String())
    fmt.Println(r.Pointer.String())

    // Output:
    // http://example.com/example.json#/foo/bar
    // /foo/bar
}
