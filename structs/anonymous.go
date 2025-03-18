package structs

import "fmt"

func DemoAnonymous() {
	fmt.Println("Anonymous Structs")

	user := struct {
		Name string
		Age  int
	}{
		Name: "Fahim Ahammed",
		Age:  28,
	}

	fmt.Println("User Name:", user.Name)
	fmt.Println("Age:", user.Age)
}
