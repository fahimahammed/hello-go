package structs

import (
	"fmt"
)

type Result struct {
	Student
	cgpa float32
}

func DemoEmbedding() {
	fmt.Println("Embedding")
	r := Result{
		Student{
			Name: "Fahim",
			Age:  23,
		},
		3.85,
	}

	fmt.Println(r)
}
