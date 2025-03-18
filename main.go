package main

import (
	"hello-go/api"
)

func main() {
	// fmt.Println("Go Structs Practice")
	// structs.DemoMethods()
	// structs.DemoPointers()
	// structs.DemoEmbedding()
	// structs.DemoAnonymous()
	// structs.DemoSlices()

	// maps.DemoMaps()

	// concurrency.DemoGoroutines()
	// concurrency.DemoChannels()

	// files.WriteToFile("example.txt", "Hello, Go!")
	// files.ReadFromFile("example.txt")

	// jsonhandler.WriteJson("person.json", jsonhandler.Person{Name: "John Doe", Age: 30, Email: "johndoe@example.com"})
	// jsonhandler.ReadJson("person.json")

	api.DemoApi()
	api.FetchUser(2)
	api.CreatePost("New Post", "This is the body of the new post.", 1)
}
