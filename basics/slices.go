// Go
package main
import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("OG:", numbers)

	numbers = append(numbers, 6)
	fmt.Println("After append:", numbers)

	numbers = numbers[1:4]
	fmt.Println("After slicing:", numbers)
}
