// Go
package main
import "fmt"

func main() {
    result := add(5, 3)
	fmt.Println("Sum:", result)
}


func add(a int, b int) int {
	return a + b
}