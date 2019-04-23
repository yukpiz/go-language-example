package main

import (
	"fmt"

	"golang.org/x/text/language"
)

func main() {
	ja := language.Japanese
	fmt.Printf("%+v\n", ja.Raw())
}
