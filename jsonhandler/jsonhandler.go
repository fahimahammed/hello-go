package jsonhandler

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person struct represents a sample JSON structure
// It contains Name, Age, and Email fields, with JSON tags for encoding/decoding
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// WriteJSON writes a Person struct as JSON to a file
func WriteJson(filename string, person Person) {

	// Convert the Person struct to a JSON-formatted byte slice with indentation
	data, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("JSON data:", string(data))

	// Write the JSON data to the specified file
	err = os.WriteFile(filename, data, 0644) // 0644 sets file permissions (read/write for owner, read-only for others)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to", filename)

}

// ReadJSON reads JSON from a file and decodes it into a Person struct
func ReadJson(filename string) {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var person Person

	err = json.Unmarshal(data, &person)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded JSON data:", person)
}
