package datatypeszerovalues
package main

import "fmt"

func main() {
    // Zero values — declared without assignment
    var i    int
    var f    float64
    var s    string
    var b    bool

    fmt.Println(i, f, s, b)
    // Output: 0  0  ""  false
    // Note: "" is empty string — it prints as nothing visible
}