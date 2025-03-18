package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// User struct represents a sample API response
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func DemoApi() {
	fmt.Println("API Demo")
}

func FetchUser(userId int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userId)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("User: %+v\n", user)
}

// CreatePost makes an API POST request
func CreatePost(title, body string, userID int) {
	postData := map[string]any{
		"title":  title,
		"body":   body,
		"userId": userID,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
