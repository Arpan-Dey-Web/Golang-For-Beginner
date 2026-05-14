package variables
package main

import "fmt"

func main() {
    var name string = "Arpan"    // string variable
    var age  int    = 22         // integer variable
    var gpa  float64 = 3.9      // float variable
    var isActive bool = true     // boolean variable
	user:= "Arpan" 				// short variable declaration (type inferred)


    fmt.Println(name, age, gpa, isActive, user)	
}	