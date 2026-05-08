package anonymousfunctions
package main

import "fmt"

func main() {
    // Defined and stored in a variable
    greet := func() {
        fmt.Println("I am an anonymous function!")
    }

    greet() // call it like a normal function
}