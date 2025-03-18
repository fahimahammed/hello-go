package structs

import "fmt"

type Student struct {
	Name string
	Age  int
}

func haveBirthday(s *Student) {
	s.Age++
}

func DemoPointers() {
	fmt.Println("Go Pointers Practice")

	s := Student{Name: "Alice", Age: 25}

	fmt.Println("Before Birthday:", s)
	haveBirthday(&s)
	fmt.Println("After Birthday:", s)
}
