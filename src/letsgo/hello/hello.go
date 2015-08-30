package main

import "fmt"
import "letsgo/stringutil"

var HELLO_WORLD = "Hello, world."

func main() {
	fmt.Println(HELLO_WORLD+"\n")
	fmt.Println(stringutil.Reverse("\n"+HELLO_WORLD))
	fmt.Println(stringutil.Reverse("\n!oG ,olleH"))
}
