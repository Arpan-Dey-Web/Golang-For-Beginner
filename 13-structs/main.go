package structs
package main

import "fmt"

// Define the struct blueprint
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Create an instance of Person
    p1 := Person{
        Name: "Arpan",
        Age:  22,
        City: "Chittagong",
    }

    // Access fields using dot notation
    fmt.Println(p1.Name) // Arpan
    fmt.Println(p1.Age)  // 22
    fmt.Println(p1)      // {Arpan 22 Chittagong}

    // Update a field
    p1.Age = 23
    fmt.Println(p1.Age)  // 23
}