# 🐹 Go (Golang) — Complete Beginner Notes

> A beginner-friendly reference covering the core fundamentals of Go.  
> Every topic includes a clear definition followed by working examples.

---

## Table of Contents

1. [Main File Structure](#1-main-file-structure)
2. [Variables](#2-variables)
3. [Constants](#3-constants)
4. [Data Types & Zero Values](#4-data-types--zero-values)
5. [fmt Package & Format Verbs (Placeholders)](#5-fmt-package--format-verbs-placeholders)
6. [Functions](#6-functions)
7. [Anonymous Functions](#7-anonymous-functions)
8. [Variable Scope](#8-variable-scope)
9. [Conditions — if / else / switch](#9-conditions--if--else--switch)
10. [Arrays](#10-arrays)
11. [Slices](#11-slices)
12. [Pointers](#12-pointers)
13. [Structs](#13-structs)

---

## 1. Main File Structure

### Definition
Every Go program starts from a **package** declaration. The special package name `main` tells Go: *"this is the entry point of the program."* The `main()` function is the first function Go runs automatically. You must import any external packages you need at the top of the file.

```go
package main   // Every executable Go program must be in package "main"

import "fmt"   // Import the fmt package so we can print to the console

func main() {
    // Your program starts executing from here
    fmt.Println("Hello, World!")
}
```

**Key Points:**
- `package main` — required for any runnable program
- `import` — loads external packages you want to use
- `func main()` — the entry point; Go runs this first, automatically
- Go files use the `.go` extension

---

## 2. Variables

### Definition
A variable is a named storage location in memory that holds a value. In Go, you must tell the compiler what type of data the variable will hold (Go is statically typed). There are two ways to declare variables.

---

### 2a. Using `var` keyword (explicit declaration)

**Syntax:** `var variableName dataType = value`

```go
package main

import "fmt"

func main() {
    var name string = "Arpan"    // string variable
    var age  int    = 22         // integer variable
    var gpa  float64 = 3.9      // float variable
    var isActive bool = true     // boolean variable

    fmt.Println(name, age, gpa, isActive)
}
```

You can also declare a variable without a value (it gets the **zero value** for its type):

```go
var name string   // zero value = ""  (empty string)
var age  int      // zero value = 0
var isOk bool     // zero value = false
```

---

### 2b. Short Variable Declaration `:=` (shorthand)

**Syntax:** `variableName := value`

Go automatically figures out the type from the value — no need to write `var` or the type name. This is the most common way to declare variables inside a function.

```go
package main

import "fmt"

func main() {
    name := "Arpan"     // Go infers this is a string
    age  := 22          // Go infers this is an int
    gpa  := 3.9         // Go infers this is a float64

    fmt.Println(name, age, gpa)
}
```

> ⚠️ **Important:** `:=` can only be used **inside a function**. For package-level variables, you must use `var`.

---

### 2c. Declaring Multiple Variables at Once

```go
// Using var
var a, b, c int = 1, 2, 3

// Using shorthand
x, y := "hello", "world"

// Block declaration (cleaner for many variables)
var (
    firstName string = "Arpan"
    lastName  string = "Dey"
    age       int    = 22
)
```

---

## 3. Constants

### Definition
A constant is a value that **cannot be changed** after it is set. Use `const` instead of `var`. Constants must be assigned a value at the time of declaration — you cannot use `:=` for constants.

```go
package main

import "fmt"

func main() {
    const pi       float64 = 3.14159
    const appName  string  = "MyApp"
    const maxRetry int     = 5

    fmt.Println(pi, appName, maxRetry)

    // pi = 3.14  ❌ This will cause a compile error — cannot change a constant
}
```

**Block constant declaration:**

```go
const (
    StatusOK    = 200
    StatusNotFound = 404
    AppVersion  = "1.0.0"
)
```

> 💡 Use constants for values that are fixed by design: math values, config keys, status codes, etc.

---

## 4. Data Types & Zero Values

### Definition
A data type tells Go what kind of data a variable holds and how much memory to allocate. Every type in Go has a **zero value** — the default value a variable gets when you declare it without assigning anything.

| Type | Description | Zero Value | Example |
|------|-------------|------------|---------|
| `int` | Whole numbers | `0` | `42`, `-10` |
| `int8` | Whole number (−128 to 127) | `0` | `100` |
| `int16` | Whole number (−32768 to 32767) | `0` | `300` |
| `int32` | Whole number | `0` | `70000` |
| `int64` | Large whole number | `0` | `9000000000` |
| `float32` | Decimal number (less precise) | `0` | `3.14` |
| `float64` | Decimal number (more precise) | `0` | `3.14159265` |
| `string` | Text / sequence of characters | `""` | `"hello"` |
| `bool` | True or False | `false` | `true`, `false` |
| `byte` | Alias for `uint8` | `0` | `65` (= 'A') |
| `rune` | Alias for `int32`, represents a Unicode character | `0` | `'A'` |

```go
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
```

---

## 5. fmt Package & Format Verbs (Placeholders)

### Definition
The `fmt` package is Go's standard library for formatted input and output. **Format verbs** (also called placeholders) are special codes that tell `fmt` how to display a value inside a string. They always start with `%`.

---

### Common fmt Functions

| Function | Description |
|----------|-------------|
| `fmt.Println()` | Prints with a newline at the end |
| `fmt.Print()` | Prints without a newline |
| `fmt.Printf()` | Prints with format verbs (no automatic newline) |
| `fmt.Sprintf()` | Returns a formatted string (doesn't print) |
| `fmt.Scan()` | Reads input from the user |

---

### Format Verbs (Placeholders)

| Verb | Meaning | Used With |
|------|---------|-----------|
| `%s` | String | `string` |
| `%d` | Integer (decimal) | `int`, `int64`, etc. |
| `%f` | Floating point | `float32`, `float64` |
| `%.2f` | Float with 2 decimal places | `float64` |
| `%t` | Boolean | `bool` |
| `%v` | Default format (works for any type) | any |
| `%T` | Type of the variable | any |
| `%p` | Pointer address | pointer |
| `\n` | Newline character | inside strings |

```go
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
```

---

## 6. Functions

### Definition
A function is a reusable block of code that performs a specific task. You define it once and can call it many times. In Go, functions are defined with the `func` keyword.

**Syntax:**
```
func functionName(parameterName parameterType) returnType {
    // code
    return value
}
```

---

### 6a. Basic Function (no parameters, no return)

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello from a function!")
}

func main() {
    greet()  // calling the function
}
```

---

### 6b. Function with Parameters

```go
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
```

---

### 6c. Function with Return Type

```go
package main

import "fmt"

// Returns an int
func add(a int, b int) int {
    return a + b
}

// Returns a string
func getGreeting(name string) string {
    return "Hello, " + name + "!"
}

func main() {
    result := add(10, 20)
    fmt.Println(result)           // 30

    message := getGreeting("Arpan")
    fmt.Println(message)          // Hello, Arpan!
}
```

---

### 6d. Function with Multiple Return Values

Go supports returning more than one value — very useful for returning a result alongside an error.

```go
package main

import "fmt"

func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Error: cannot divide by zero"
    }
    return a / b, "success"
}

func main() {
    result, status := divide(10, 2)
    fmt.Println(result, status)   // 5  success

    result2, status2 := divide(10, 0)
    fmt.Println(result2, status2) // 0  Error: cannot divide by zero
}
```

---

### 6e. Named Return Values

You can name the return variables in the function signature and use a bare `return`.

```go
package main

import "fmt"

func minMax(a, b int) (min int, max int) {
    if a < b {
        min = a
        max = b
    } else {
        min = b
        max = a
    }
    return // bare return — returns min and max
}

func main() {
    lo, hi := minMax(15, 7)
    fmt.Println("Min:", lo, "Max:", hi) // Min: 7 Max: 15
}
```

---

### 6f. Storing a Function in a Variable

Functions in Go are **first-class values** — you can store them in a variable and call them through that variable.

```go
package main

import "fmt"

func main() {
    // Store a function in a variable
    sayHello := func(name string) {
        fmt.Println("Hello,", name)
    }

    sayHello("Arpan")  // Output: Hello, Arpan

    // You can also store a named function
    add := func(a, b int) int {
        return a + b
    }

    result := add(5, 3)
    fmt.Println(result) // 8
}
```

---

## 7. Anonymous Functions

### Definition
An anonymous function is a function that has **no name**. It is defined inline and is often used when you need a quick, one-time-use function. In Go, anonymous functions can be assigned to variables, passed as arguments, or invoked immediately.

---

### 7a. Basic Anonymous Function

```go
package main

import "fmt"

func main() {
    // Defined and stored in a variable
    greet := func() {
        fmt.Println("I am an anonymous function!")
    }

    greet() // call it like a normal function
}
```

---

### 7b. Immediately Invoked Anonymous Function (IIFE)

You can define and call an anonymous function in the same line by adding `()` at the end.

```go
package main

import "fmt"

func main() {
    // Defined AND immediately called
    func() {
        fmt.Println("This runs immediately!")
    }() // <-- the () here invokes it right away

    // With arguments
    func(x, y int) {
        fmt.Println("Sum:", x+y)
    }(10, 20) // passes 10 and 20 as arguments
}
```

---

### 7c. Anonymous Function as a Function Argument

```go
package main

import "fmt"

// This function takes another function as a parameter
func doMath(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    result := doMath(10, 5, func(a, b int) int {
        return a * b
    })
    fmt.Println("Result:", result) // 50
}
```

---

## 8. Variable Scope

### Definition
Scope defines **where in your code a variable is accessible**. A variable can only be used within the block `{ }` it was declared in. Go has three main levels of scope:

1. **Package scope** — declared outside all functions; accessible anywhere in the file (and package)
2. **Function scope** — declared inside a function; accessible only within that function
3. **Block scope** — declared inside an `if`, `for`, or other block; accessible only inside that block

```go
package main

import "fmt"

var packageLevel = "I am accessible everywhere in this package" // package scope

func main() {
    functionLevel := "I am only accessible inside main()" // function scope

    if true {
        blockLevel := "I only exist inside this if block" // block scope
        fmt.Println(blockLevel)
    }

    // fmt.Println(blockLevel) // ❌ ERROR: blockLevel is not accessible here

    fmt.Println(packageLevel)
    fmt.Println(functionLevel)
}
```

---

## 9. Conditions — if / else / switch

### Definition
Conditions let your program make decisions. Go has `if/else` for true/false branching and `switch` for matching a value against multiple cases.

---

### 9a. if / else

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else if age >= 13 {
        fmt.Println("You are a teenager.")
    } else {
        fmt.Println("You are a child.")
    }
}
```

---

### 9b. Variable Declared Inside if Block

Go allows you to declare a variable directly in the `if` statement. That variable only exists inside the `if/else` block.

```go
package main

import "fmt"

func main() {
    // 'score' is declared here and only lives inside this if/else
    if score := 85; score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }

    // fmt.Println(score) // ❌ ERROR — score is out of scope here
}
```

---

### 9c. Comparison Operators

| Operator | Meaning |
|----------|---------|
| `==` | Equal to |
| `!=` | Not equal to |
| `>` | Greater than |
| `<` | Less than |
| `>=` | Greater than or equal |
| `<=` | Less than or equal |

### 9d. Logical Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `&&` | AND — both must be true | `age > 18 && hasID` |
| `\|\|` | OR — at least one must be true | `isAdmin \|\| isMod` |
| `!` | NOT — inverts true/false | `!isLoggedIn` |

---

### 9e. switch / case

`switch` compares a value against several possible cases and runs the matching one. No `break` needed — Go exits automatically after the matching case.

```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("Start of the work week.")
    case "Friday":
        fmt.Println("Almost the weekend!")
    case "Saturday", "Sunday":
        fmt.Println("It is the weekend!")
    default:
        fmt.Println("It is a regular weekday.")
    }
}
```

**Switch without a value** (acts like if/else):

```go
package main

import "fmt"

func main() {
    score := 75

    switch {
    case score >= 90:
        fmt.Println("Excellent")
    case score >= 70:
        fmt.Println("Good")
    default:
        fmt.Println("Needs Improvement")
    }
}
```

---

## 10. Arrays

### Definition
An array is a **fixed-size, ordered collection of elements of the same type**. The size of an array is part of its type — once set, it cannot change.

**Syntax:** `var arrayName [size]dataType`

```go
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
```

> ⚠️ Arrays in Go are **value types** — when you assign an array to another variable, a full copy is made.

---

## 11. Slices

### Definition
A slice is a **flexible, dynamic view into an array**. Unlike arrays, slices have no fixed size — they can grow and shrink. Slices are built on top of an underlying array and are the most commonly used collection in Go.

---

### 11a. Declaring a Slice

```go
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
```

---

### 11b. Length and Capacity

- **Length (`len`)** — the number of elements currently in the slice
- **Capacity (`cap`)** — how many elements the underlying array can hold from the slice's starting position

```go
package main

import "fmt"

func main() {
    s := make([]int, 3, 6)
    fmt.Println("Length:", len(s))   // 3
    fmt.Println("Capacity:", cap(s)) // 6
}
```

---

### 11c. Appending to a Slice

Use `append()` to add elements. If the slice runs out of capacity, Go automatically creates a new, larger underlying array.

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3}
    nums = append(nums, 4, 5)
    fmt.Println(nums) // [1 2 3 4 5]

    // Append another slice using ...
    extra := []int{6, 7}
    nums = append(nums, extra...)
    fmt.Println(nums) // [1 2 3 4 5 6 7]
}
```

---

### 11d. How Slices Work Internally (Pointer, Length, Capacity)

Internally, a slice is a small struct with three fields:
1. **Pointer** — points to the first element of the slice in the underlying array
2. **Length** — how many elements are in the slice
3. **Capacity** — how many elements the underlying array has from that pointer onward

```go
package main

import "fmt"

func main() {
    arr := [6]int{10, 20, 30, 40, 50, 60}

    s1 := arr[1:4]  // starts at index 1, ends before 4
    // Pointer → arr[1] (value: 20)
    // Length  → 3 (elements: 20, 30, 40)
    // Capacity → 5 (from index 1 to end of arr)

    fmt.Println("s1:", s1)          // [20 30 40]
    fmt.Println("len:", len(s1))    // 3
    fmt.Println("cap:", cap(s1))    // 5
}
```

> ⚠️ Because a slice shares memory with its underlying array, **modifying a slice also modifies the original array.**

---

## 12. Pointers

### Definition
A pointer is a variable that **stores the memory address** of another variable — not the value itself, but the location where that value lives in memory.

---

### 12a. Pointer Operators

| Operator | Name | Description |
|----------|------|-------------|
| `&` | Address-of | Gets the memory address of a variable |
| `*` | Dereference | Accesses the value at the memory address |

```go
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
```

---

### 12b. Pass by Value vs Pass by Reference

**Pass by Value** — a copy of the value is sent to the function. Changes inside the function do NOT affect the original.

```go
package main

import "fmt"

func addTen(num int) {
    num = num + 10       // only changes the local copy
    fmt.Println("Inside function:", num) // 30
}

func main() {
    x := 20
    addTen(x)
    fmt.Println("Outside function:", x) // still 20 — original not changed
}
```

**Pass by Reference** — a pointer (memory address) is sent to the function. Changes inside the function DO affect the original.

```go
package main

import "fmt"

func addTen(num *int) {       // parameter type is *int (pointer to int)
    *num = *num + 10          // dereference and update the original value
    fmt.Println("Inside function:", *num) // 30
}

func main() {
    x := 20
    addTen(&x)                // pass the address of x
    fmt.Println("Outside function:", x)  // 30 — original was changed!
}
```

---

### 12c. Pointer Return Type

A function can return a pointer, so the caller can access and modify the original memory.

```go
package main

import "fmt"

func newValue() *int {
    val := 100
    return &val   // return the address of val
}

func main() {
    ptr := newValue()
    fmt.Println("Value:", *ptr) // 100

    *ptr = 999
    fmt.Println("Updated:", *ptr) // 999
}
```

---

## 13. Structs

### Definition
A struct is a **user-defined type that groups related fields together**. Think of it like a blueprint or template that defines what data a certain "thing" should have. Each field has a name and a type.

---

### 13a. Defining and Using a Struct

```go
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
```

---

### 13b. Nested Struct

A struct can contain another struct as one of its fields.

```go
package main

import "fmt"

type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // nested struct
}

func main() {
    p := Person{
        Name: "Arpan",
        Age:  22,
        Address: Address{
            City:    "Chittagong",
            Country: "Bangladesh",
        },
    }

    fmt.Println(p.Name)            // Arpan
    fmt.Println(p.Address.City)    // Chittagong
    fmt.Println(p.Address.Country) // Bangladesh
}
```

---

### 13c. Constructor Function

Go does not have a built-in constructor like Java or C++. Instead, the convention is to create a regular function — usually named `NewTypeName` — that creates and returns a new instance of the struct.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Constructor function
func NewPerson(name string, age int) Person {
    return Person{
        Name: name,
        Age:  age,
    }
}

// Constructor returning a pointer (more common in real apps)
func NewPersonPtr(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

func main() {
    p1 := NewPerson("Arpan", 22)
    fmt.Println(p1) // {Arpan 22}

    p2 := NewPersonPtr("Dey", 25)
    fmt.Println(p2)  // &{Dey 25}
    fmt.Println(*p2) // {Dey 25}
}
```

---

### 13d. Methods on a Struct (Functions Defined on a Type)

### Definition
A method is a function that **belongs to a specific struct type**. It has a special parameter before the function name called the **receiver** — this is what connects the function to the struct. Using methods lets you call functions directly on a struct instance using dot notation.

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method with a VALUE receiver — gets a copy, cannot modify original
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with a VALUE receiver
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Method with a POINTER receiver — can modify the original struct
func (r *Rectangle) Scale(factor float64) {
    r.Width  = r.Width * factor
    r.Height = r.Height * factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}

    fmt.Println("Area:", rect.Area())           // 50
    fmt.Println("Perimeter:", rect.Perimeter()) // 30

    rect.Scale(2)
    fmt.Println("After Scale:", rect.Width, rect.Height) // 20 10
}
```

**Value receiver vs Pointer receiver:**

| | Value Receiver `(r Rectangle)` | Pointer Receiver `(r *Rectangle)` |
|--|--|--|
| Gets | A copy of the struct | A pointer to the original struct |
| Can modify original? | ❌ No | ✅ Yes |
| Use when | Just reading fields | Modifying fields |

---

### 13e. Struct with Function-Type Fields

A struct field can itself be a function. This lets different instances have different behavior stored inside the same struct shape.

```go
package main

import "fmt"

type Animal struct {
    Name  string
    Sound func() string   // this field holds a function
}

func main() {
    dog := Animal{
        Name: "Dog",
        Sound: func() string {
            return "Woof!"
        },
    }

    cat := Animal{
        Name: "Cat",
        Sound: func() string {
            return "Meow!"
        },
    }

    fmt.Println(dog.Name, "says:", dog.Sound()) // Dog says: Woof!
    fmt.Println(cat.Name, "says:", cat.Sound()) // Cat says: Meow!
}
```

---

### 13f. Struct Comparison and Anonymous Struct

```go
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    // Structs with same values are equal
    p1 := Point{1, 2}
    p2 := Point{1, 2}
    fmt.Println(p1 == p2) // true

    // Anonymous struct — useful for one-off data grouping
    config := struct {
        Host string
        Port int
    }{
        Host: "localhost",
        Port: 8080,
    }
    fmt.Println(config.Host, config.Port)
}
```

---

## Things to Remember

| Topic | Key Rule |
|-------|----------|
| `:=` | Shorthand only inside functions |
| Constants | Cannot be changed after declaration |
| Zero values | Every type has a default: `0`, `""`, `false` |
| Arrays | Fixed size, value type (copy on assignment) |
| Slices | Dynamic, reference type (shares memory with array) |
| Pointers `&` | Gets the address |
| Pointers `*` | Gets the value at the address |
| Pass by value | Function gets a copy — original safe |
| Pass by reference | Function gets address — original can change |
| Struct methods | Value receiver = read only; Pointer receiver = can modify |

---

> **GitHub:** [github.com/Arpan-Dey-Web](https://github.com/Arpan-Dey-Web)  
> **Portfolio:** [arpandeyweb.vercel.app](https://arpandeyweb.vercel.app)