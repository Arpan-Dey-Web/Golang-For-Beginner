package pointers
package main

import "fmt"

func main() {
    age := 22

    ptr := &age   // ptr now holds the memory address of age
    fmt.Println("Value of age:", age)       // 22
    fmt.Println("Address of age:", &age)    // 0xc000014090 (some memory address)
    fmt.Println("ptr holds:", ptr)          // same address as &age
    fmt.Println("Value via ptr:", *ptr)     // 22 — dereferencing the pointer

    // Changing the value through the pointer
    *ptr = 25
    fmt.Println("age after change:", age)   // 25 — original variable changed!
}