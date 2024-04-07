package main

import "fmt"

func hello(str string) string {
    return "Hello, " + str + "!" 
}

func main() {
    name := "Evercode"
    greeting := hello(name)
    fmt.Println(greeting)
}