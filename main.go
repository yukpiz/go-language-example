package main

import (
	"fmt"

	"golang.org/x/text/language"
)

func main() {
	ja := language.Japanese
	en := language.AmericanEnglish
	fmt.Printf("%+v\n", ja)
	fmt.Printf("%+v\n", en)
	fmt.Printf("%d\n", ja)
	fmt.Printf("%d\n", en)
}
