package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() {
	fmt.Println("area:", 0.5*t.base*t.height)
}

// Method yg memiliki return
func (t triangle) luas() float64 {
	return 0.5 * t.base * t.height
}

func (t *triangle) increaseSize() {
	t.base += 1
	t.height += 1
}

func main() {
	instanceTriangle := triangle{10, 12}
	instanceTriangle.area()

	luasSegitiga := triangle{24, 30}
	luas := luasSegitiga.luas()
	fmt.Println("luas:", luas)

	fmt.Println("sebelum:", instanceTriangle)
	instanceTriangle.increaseSize()
	fmt.Println("sesudah:", instanceTriangle)
}
