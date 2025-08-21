// Go
package main
import "fmt"

type USER struct {
	Name string
	Age  int
}

func main() {
	user := USER{
		Name: "John Doe",
		Age:  30,
	}
	fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)

	userPtr := &user
	userPtr.Age = 31
	fmt.Printf("Updated Age: %d\n", userPtr.Age)
	fmt.Printf("User Struct: %+v\n", user)
	fmt.Printf("User Struct (Pointer): %+v\n", userPtr)

	userCopy := user
	userCopy.Name = "Jane Doe"
	fmt.Printf("Original User Name: %s, Copied User Name: %s\n", user.Name, userCopy.Name)
	fmt.Printf("User Struct (Copy): %+v\n", userCopy)
	fmt.Printf("User Struct (Pointer Copy): %+v\n", &userCopy)
	fmt.Printf("User Struct (Pointer Original): %+v\n", &user)
}
