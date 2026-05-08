package slices
package main

import "fmt"

func main() {
    // Method 1: Declare a slice literal (most common)
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println(fruits) // [apple banana cherry]

    // Method 2: Using make(type, length, capacity)
    numbers := make([]int, 3, 5)
    fmt.Println(numbers) // [0 0 0]

    // Method 3: Slice from an array
    arr := [5]int{10, 20, 30, 40, 50}
    slice := arr[1:4]   // includes index 1, 2, 3 — excludes index 4
    fmt.Println(slice)  // [20 30 40]
}