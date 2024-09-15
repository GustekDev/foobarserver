package main

import (
	"fmt"

	"github.com/GustekDev/foobarserver/internal/foo"
)

func main() {
	x := foo.Foobar()
	fmt.Println(x + " from the server")
}
