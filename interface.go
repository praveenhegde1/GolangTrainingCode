package main
import (
	"fmt"
	"math"
)
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}


func (c circle) Area() float64 { 
	return math.Pi * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius

	
}
func measure(s Shape){
fmt.Println(s.Area())
fmt.Println(s.Perimeter())
fmt.Println(s)
}
func main() {
	rect := Rectangle{width: 5, height: 3}
	circ := circle{radius : 10}
	measure(rect)
	measure(circ)
	
}
