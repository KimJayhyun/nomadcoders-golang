package main

import (
	"fmt"

	"github.com/kimjayhyun/study-golang/mydict"
)

func main() {
    dictionary := mydict.Dictionary{"test": "this is just a test"}

    definition, err := dictionary.Search("test")

    if err != nil {
        fmt.Println(err)
    }

    errAdd := dictionary.Add("name", "jhkim")

    if errAdd != nil {
        fmt.Println(errAdd)
    }

    fmt.Println(definition)
}
