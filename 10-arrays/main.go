package arrays
package main

import "fmt"

func main() {
    // Declare an array of 5 integers
    var scores [5]int
    fmt.Println(scores) // [0 0 0 0 0] — zero values

    // Declare and initialize
    fruits := [3]string{"apple", "banana", "cherry"}
    fmt.Println(fruits) // [apple banana cherry]

    // Let Go count the size automatically using [...]
    colors := [...]string{"red", "green", "blue"}
    fmt.Println(colors)        // [red green blue]
    fmt.Println(len(colors))   // 3

    // Access elements by index (starts at 0)
    fmt.Println(fruits[0])  // apple
    fmt.Println(fruits[2])  // cherry

    // Update an element
    fruits[1] = "mango"
    fmt.Println(fruits) // [apple mango cherry]

    // Loop through an array
    for i := 0; i < len(fruits); i++ {
        fmt.Println(i, fruits[i])
    }

    // Using range (cleaner loop)
    for index, value := range fruits {
        fmt.Println(index, value)
    }
}