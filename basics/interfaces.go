// Go
package main
import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
    s := Circle{Radius: 5}
	fmt.Printf("Area of the circle: %.2f\n", s.Area())
}
