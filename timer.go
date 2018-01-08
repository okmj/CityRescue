package main

import (
	"fmt"

	"github.com/gorilla/securecookie"
)

var hashKey []byte
var blockKey []byte

func test() {
	fmt.Println("hashkey:", securecookie.GenerateRandomKey(24))
	fmt.Println("blockKey:", securecookie.GenerateRandomKey(24))

}
