package main

import "fmt"

var packageLevel = "I am accessible everywhere in this package" // package scope

func main() {
    functionLevel := "I am only accessible inside main()" // function scope

    if true {
        blockLevel := "I only exist inside this if block" // block scope
        fmt.Println(blockLevel)
    }

    // fmt.Println(blockLevel) // ❌ ERROR: blockLevel is not accessible here

    fmt.Println(packageLevel)
    fmt.Println(functionLevel)
}