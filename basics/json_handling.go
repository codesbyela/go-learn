// Go
package main
import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
    user := User{
		Name:  "Alice",
		Age:   30,
		Email: "q7oMl@example.com",
	}

	jsonData, _ := json.Marshal(user)
	fmt.Println("JSON Data:", string(jsonData))

	jsonString := `{"name":"Bob","age":25,"email":"m2Mw3@example.com"}`
	var userFromJSON User
	err := json.Unmarshal([]byte(jsonString), &userFromJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("User from JSON: %+v\n", userFromJSON)	
}
