package intermediate

import "fmt"

const PI = 3.14

type geometry interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return PI * c.radius * c.radius
}

func getArea(g geometry) {
	fmt.Println("The area of this shape is", g.area())
}

func main() {
	r := rectangle{width: 5, height: 4}
	c := circle{radius: 4}

	getArea(r)
	getArea(c)
}
