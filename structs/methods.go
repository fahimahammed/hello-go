package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func DemoMethods() {
	p := Person{Name: "Alice", Age: 25}
	p.Display()
	p.HaveBirthday()
	fmt.Println("After Birthday:")
	p.Display()
}
