// Go
package main
import (
	"fmt"
	"os"
)

func main() {
	// Write file
    data := []byte("Hello, Go file I/O!")
	err := os.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully!")

	// Read file
	content, err := os.ReadFile("example.txt")
	if err != nil {		
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("File content:", string(content))


	//UPDATE FILE CONTENT
	err = os.WriteFile("example.txt", []byte("Updated content"), 0644)
	if err != nil {		
		fmt.Println("Error updating file:", err)
		return
	}
	content, err = os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading updated file:", err)
		return
	}
	fmt.Println("Updated file content:", string(content))

	// DELETE FILE
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully!")
}
