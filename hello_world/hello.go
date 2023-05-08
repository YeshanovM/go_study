package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println("Hello, world!\nHere's some smart quote I think I agree with:")
    fmt.Println("\"" + quote.Opt() + "\"")
}