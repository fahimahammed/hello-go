# GoLang

## **Introduction to Go**
### **What is Go?**  
Go (or Golang) is an open-source programming language developed by Google in 2007 and released in 2009. It is designed for simplicity, performance, and concurrency, making it ideal for building scalable and efficient software. Go is statically typed, compiled, and has garbage collection, making it a powerful yet easy-to-use language.  

#### **Key Features of Go:**
- **Concurrency**: Built-in support for goroutines and channels.
- **Fast Compilation**: Compiled language with quick compilation times.
- **Garbage Collection**: Automatic memory management.
- **Static Typing**: Ensures type safety while maintaining simplicity.
- **Cross-Platform**: Can be compiled to run on different operating systems.
- **Standard Library**: Comes with a rich standard library for networking, web development, and more.
- **Minimalistic Syntax**: Inspired by C but with a focus on readability and simplicity.

---

### **Comparison Between Go and Other Languages**
| Feature         | Go (Golang) | Python | Java | C++ | JavaScript |
|---------------|------------|--------|------|-----|------------|
| **Type System** | Statically typed | Dynamically typed | Statically typed | Statically typed | Dynamically typed |
| **Compilation** | Compiled | Interpreted | Compiled (JVM) | Compiled | Interpreted |
| **Performance** | Fast (compiled) | Slower than Go | Slower than Go but optimized | Very fast | Slower (runtime execution) |
| **Concurrency** | Built-in goroutines and channels | Thread-based, needs multiprocessing | Thread-based (heavy threads) | Thread-based | Asynchronous (Event Loop) |
| **Memory Management** | Garbage collection | Garbage collection | Garbage collection | Manual or smart pointers | Garbage collection |
| **Ease of Learning** | Simple and minimalistic | Easy | Moderate | Complex | Easy |
| **Use Cases** | Cloud computing, microservices, web backend, networking | AI, ML, scripting, web | Enterprise apps, Android development | System programming, game development | Web development, front-end |

---

### **When to Use Go?**
- **Building Microservices & Cloud Applications**: Used by companies like Uber, Netflix, and Dropbox.
- **High-Performance Web Servers**: Frameworks like Gin and Fiber make Go great for APIs.
- **Networking & System Programming**: Excellent for handling concurrent network connections.
- **DevOps & CLI Tools**: Tools like Kubernetes, Docker, and Terraform are built with Go.

---

### **How Go Works (For Beginners)**  

✅ **Write & Run Code** – Save in `.go` file, run with `go run file.go`.  
✅ **Fast & Simple** – Compiles to machine code for speed.  
✅ **Variables & Types** – Use `:=` for type inference.  
✅ **Functions** – Reusable code blocks.  
✅ **Goroutines** – Run tasks in parallel using `go` keyword.  
✅ **Channels** – Communicate between tasks with `make(chan type)`.  
✅ **Built-in Tools** – `go run`, `go build`, `go fmt`, `go test`.  

---

 
### **Go Language Syntax**  

---

#### **1. Hello World**
The first program in any language is printing "Hello, World!" to the console.  

##### **Example:**  
```go
package main  // Defines the package

import "fmt"  // Imports the fmt package for formatted I/O

func main() {
    fmt.Println("Hello, World!")  // Prints output to the console
}
```
##### **Explanation:**
- `package main`: Every Go program must have a `main` package.
- `import "fmt"`: Imports the `fmt` package for formatted output.
- `func main() {}`: Defines the main function, which is the program's entry point.
- `fmt.Println()`: Prints a string followed by a new line.

---

#### **2. Data Types**
Go is a statically typed language, meaning variables must have a specific type.

##### **Integer Data Types in Go**
| **Type**   | **Size**  | **Description** |
|------------|---------|----------------|
| `int`      | Depends on system (32-bit/64-bit) | Default integer type |
| `int8`     | 8-bit   | Stores values from -128 to 127 |
| `int16`    | 16-bit  | Stores values from -32,768 to 32,767 |
| `int32`    | 32-bit  | Stores values from -2,147,483,648 to 2,147,483,647 |
| `int64`    | 64-bit  | Stores values from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `uint8`    | 8-bit   | Stores values from 0 to 255 |
| `uint16`   | 16-bit  | Stores values from 0 to 65,535 |
| `uint32`   | 32-bit  | Stores values from 0 to 4,294,967,295 |
| `uint64`   | 64-bit  | Stores values from 0 to 18,446,744,073,709,551,615 |

##### **Floating Point Data Types**
| **Type**   | **Size**  | **Description** |
|------------|---------|----------------|
| `float32`  | 32-bit  | 6-7 decimal places precision |
| `float64`  | 64-bit  | 15-16 decimal places precision |

##### **Boolean Data Type**
| **Type**  | **Values**  |
|----------|-------------|
| `bool`   | `true` or `false` |

##### **String Data Type**
| **Type**  | **Description** |
|----------|----------------|
| `string` | Stores text data |

##### **Example:**
```go
package main

import "fmt"

func main() {
    var a int = 10
    var b int64 = 1000000000
    var c float64 = 20.5
    var d bool = true
    var e string = "GoLang"

    fmt.Println(a, b, c, d, e)
}
```

---

#### **3. Keywords**
Go has **25 reserved keywords**, which cannot be used as variable names.  

| **Keywords**        | **Description** |
|--------------------|--------------|
| `break`, `continue` | Loop control statements |
| `if`, `else`, `switch`, `case` | Control flow |
| `for`, `range` | Loops |
| `var`, `const`, `type` | Variable and constant declarations |
| `func`, `return`, `defer` | Function declarations and behavior |
| `package`, `import` | Package and import management |
| `go`, `select`, `chan`, `map` | Concurrency and collections |
| `interface`, `struct` | Data structures |
| `default`, `fallthrough` | Switch-case behavior |

---

#### **4. Tokens - Escape Sequences & Comments**
##### **Escape Sequences**
Escape sequences are special characters used in strings.
| **Escape Sequence** | **Description** |
|---------------------|----------------|
| `\n`               | New line |
| `\t`               | Tab space |
| `\"`               | Double quote |
| `\\`               | Backslash |

##### **Example:**  
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello\tGo\nNew Line") // \t (Tab), \n (New Line)
}
```

##### **Comments**
```go
// Single-line comment

/*
Multi-line comment
*/
```

---

#### **5. Variables**
Variables store values in memory.

##### **Example:**  
```go
package main

import "fmt"

func main() {
    var name string = "Alice" // Explicit type
    age := 25  // Short variable declaration (Type inferred)

    fmt.Println(name, age)
}
```

---

#### **6. Formatting Output**
Go provides `fmt.Printf()` for formatted output.

##### **Example:**
```go
package main

import "fmt"

func main() {
    name := "Bob"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age) // %s for string, %d for integer
}
```

---

#### **7. Constants**
Constants are fixed values that do not change.

##### **Example:**
```go
package main

import "fmt"

const Pi = 3.1416

func main() {
    fmt.Println("Value of Pi:", Pi)
}
```

---

#### **8. Getting User Input**
The `fmt.Scanln()` function reads user input.

##### **Example:**
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name) // Read input
    fmt.Println("Hello,", name)
}
```

---

#### **9. Number Formatting**
##### **Example:**
```go
package main

import "fmt"

func main() {
    num := 255
    fmt.Printf("Decimal: %d\nHex: %x\nBinary: %b\n", num, num, num)
}
```

---

#### **10. Operators**
| **Operator** | **Description** |
|-------------|----------------|
| `+` `-` `*` `/` `%` | Arithmetic operators |
| `==` `!=` `<` `>` `<=` `>=` | Comparison operators |
| `&&` `||` `!` | Logical operators |

##### **Example:**
```go
package main

import "fmt"

func main() {
    a, b := 10, 3
    fmt.Println("Addition:", a+b)
}
```

---

#### **11. If, Else If, Else**
##### **Example:**
```go
if num > 0 {
    fmt.Println("Positive")
} else {
    fmt.Println("Negative")
}
```

---

#### **12. Switch-Case**
##### **Example:**
```go
switch day {
case 1:
    fmt.Println("Monday")
default:
    fmt.Println("Other Day")
}
```

---

#### **13. Loop Control Statements**
##### **Example:**
```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

---

#### **14. Break and Continue**
##### **Example:**
```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)
}
```

---

#### **15. Functions**
##### **Example:**
```go
func add(a, b int) int {
    return a + b
}
```

---

#### **16. Structure (Detailed)**
A **struct** is a collection of fields used to group related data together.

##### **Defining a Struct**
```go
type Person struct {
    Name string
    Age  int
}
```

##### **Using Structs**
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println("Name:", p.Name, "Age:", p.Age)
}
```

##### **Struct with Methods**
```go
func (p Person) greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```
---

#### **17. Pointers (Detailed)**
A pointer is a variable that stores the memory address of another variable.

##### **Why Use Pointers?**
- To avoid copying large data structures.
- To modify values inside functions.
- Efficient memory management.

#### **Declaring a Pointer**
```go
var ptr *int // Pointer to an integer
```

#### **Example:**
```go
package main

import "fmt"

func main() {
    var x int = 10
    var p *int = &x // Pointer to x

    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Value stored at pointer p:", *p) // Dereferencing
}
```

#### **Passing Pointers to Functions**
```go
func updateValue(ptr *int) {
    *ptr = *ptr + 10
}

func main() {
    num := 20
    updateValue(&num)
    fmt.Println("Updated Value:", num)
}
```
---

#### **18. Arrays & Slices (Detailed)**
##### **Arrays in Go**
- Fixed size.
- Cannot be resized.

```go
var arr [3]int = [3]int{1, 2, 3}
fmt.Println(arr[0]) // Accessing element
```

##### **Slices in Go**
- Dynamic arrays (resizable).
- Uses `append()` function.

```go
slice := []int{4, 5, 6}
slice = append(slice, 7) // Appending elements
```

---

#### **19. Maps (Key-Value Pairs)**
- **Maps** are unordered collections of key-value pairs.

##### **Creating a Map**
```go
user := map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(user["Alice"]) 
```

##### **Adding & Deleting Entries**
```go
user["Charlie"] = 35 // Add
delete(user, "Bob")  // Delete
```

---

#### **20. Goroutines (Concurrency)**
- Goroutines are lightweight threads for concurrent execution.
- Uses `go` keyword.

##### **Example:**
```go
package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 5; i++ {
        fmt.Println(msg)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    go printMessage("Hello")
    go printMessage("World")

    time.Sleep(time.Second * 3) // Give time for goroutines to execute
}
```

##### **Use Case:**
- Running background tasks like **HTTP requests**.
- Parallel processing for faster execution.
- Avoiding blocking operations.

---

