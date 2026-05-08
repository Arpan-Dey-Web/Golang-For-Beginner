package fmtpackageformatverbs
package main

import "fmt"

func main() {
    name    := "Arpan"
    age     := 22
    gpa     := 3.95
    isActive := true

    // %s for string
    fmt.Printf("Name: %s\n", name)

    // %d for integer
    fmt.Printf("Age: %d\n", age)

    // %f for float (6 decimal places by default)
    fmt.Printf("GPA: %f\n", gpa)

    // %.2f for float with exactly 2 decimal places
    fmt.Printf("GPA (rounded): %.2f\n", gpa)

    // %t for boolean
    fmt.Printf("Active: %t\n", isActive)

    // %v prints any variable in its default format
    fmt.Printf("Name: %v, Age: %v\n", name, age)

    // %T prints the TYPE of the variable
    fmt.Printf("Type of age: %T\n", age)       // int
    fmt.Printf("Type of gpa: %T\n", gpa)       // float64

    // Sprintf — returns formatted string without printing
    message := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
    fmt.Println(message)
}