package structs

import "fmt"

func DemoSlices() {
	fmt.Println("Go Slices Practice")
	users := []Person{
		{
			Name: "Fahim Ahammed Firoz",
			Age:  34,
		},
		{
			Name: "Rahima Begum",
			Age:  32,
		},
	}

	fmt.Println(users)

	// old style
	for i := 0; i < len(users); i++ {
		users[i].Display()
	}

	for i := range len(users) { // More idiomatic way
		users[i].Display()
	}

	for _, user := range users {
		user.Display()
	}
}
