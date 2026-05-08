package constants
package main

import "fmt"

func main() {
    const pi       float64 = 3.14159
    const appName  string  = "MyApp"
    const maxRetry int     = 5

    fmt.Println(pi, appName, maxRetry)

    // pi = 3.14  ❌ This will cause a compile error — cannot change a constant
}