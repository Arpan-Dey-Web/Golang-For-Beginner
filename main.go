package main

import "fmt"

// import "fmt"

// func makeCoffee(kind string) string {
// 	fmt.Println("Making coffee...")
// 	return "Coffee made: " + kind
// }

// func main() {
// 	// Variable declaration and initialization
// 	// var name string = "Arpan"
// 	name := "Arpan"
// 	age := 20
// 	var studentName2 = "Arpan"

// 	var (
// 		name1        string = "Arpan1"
// 		age1                = "21"
// 		studentName1        = "Arpan2"
// 	)

// 	fmt.Println("Hello, " + name + "!")
// 	fmt.Println(age)
// 	fmt.Println(age1)
// 	fmt.Println("Hello, " + name1 + "!")
// 	fmt.Println("Hello, " + studentName1 + "!")
// 	fmt.Println("Hello, " + studentName2 + "!")

// 	var a, b, c int = 1, 2, 3
// 	fmt.Println(a, b, c)

// 	fmt.Println(makeCoffee("Black "))

// }

// func makeCOffee( ){
// 	fmt.Println("Making Coffee......")
// }

func main() {
	// Function literal
	// makeCoffee:=func (){
	// 	fmt.Println("Making Coffee......")
	// }
	// makeCoffee()

	// variable declare with in if block and also check the condition with that variable
	// if score := 85; score >= 90 {
	// 	fmt.Println("Grade: A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade: B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade: C")
	// } else {
	// 	fmt.Println("Grade: F")
	// }


	a:= 500
	b:=&a
	c:=*b

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b (address of a):", b)
	fmt.Println("Value of c (value at address b):", c)
	fmt.Println("Address of b:", &b)
	fmt.Println("Address of c:", &c)
	fmt.Println("Address of c:", c)

}
