package maps

import "fmt"

func DemoMaps() {
	fmt.Println("Go Maps Practice")
	studentResults := map[string]int{
		"Alice":   90,
		"Bob":     80,
		"Charlie": 70,
	}

	fmt.Println("Student Results:", studentResults)

	delete(studentResults, "Bob")
	fmt.Println("Student Results after deleting Bob:", studentResults)

	studentResults["David"] = 95
	fmt.Println("Student Results after adding David:", studentResults)

	for name, score := range studentResults {
		fmt.Printf("%s: %d\n", name, score)
	}
}
