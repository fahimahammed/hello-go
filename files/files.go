package files

import (
	"fmt"
	"os"
)

func WriteToFile(filename, content string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content written to file successfully.")
}

func ReadFromFile(filename string) {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File contents:", string(data))
}
