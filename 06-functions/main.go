package functions
package main

import "fmt"

// Takes two int parameters and prints their sum
func add(a int, b int) {
    fmt.Println("Sum:", a+b)
}

// Shorthand: if both params are same type
func multiply(a, b int) {
    fmt.Println("Product:", a*b)
}

func main() {
    add(10, 20)       // Output: Sum: 30
    multiply(4, 5)    // Output: Product: 20
}