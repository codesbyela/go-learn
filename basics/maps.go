// Go
package main
import "fmt"

func main() {
    scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Charlie": 95,
	}

	scores["David"] = 80 // Adding a new key-value pair
	fmt.Println("Scores:", scores)

	delete(scores, "Bob") // Deleting a key-value pair
	fmt.Println("Scores after deletion:", scores)

	fmt.Println("Alice's score:", scores["Alice"]) // Accessing a value by key

	value, exists := scores["Charlie"] // Checking if a key exists
	if exists {
		fmt.Println("Charlie's score:", value)
	} else {
		fmt.Println("Charlie's score not found")
	}
	fmt.Println("Total number of students:", len(scores)) // Getting the number of key-value pairs
}
