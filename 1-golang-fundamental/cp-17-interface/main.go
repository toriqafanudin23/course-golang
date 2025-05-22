package main

import "fmt"

// Menyimpan method yang belum terpakai
type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

type Square struct {
	side float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.length
}

func (r Rectangle) getPerimeter() float64 {
	return 2 * (r.width + r.length)
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

func getAreaOfRectangle(r Rectangle) {
	fmt.Println("Luas:", r.getArea())
}

func getAreaOfSquare(s Square) {
	fmt.Println("Luas:", s.getArea())
}

// Menggabungkan getAreaOfRectangle dan getAreaOfSquare
func getArea(s Shape) {
	fmt.Println("Luas:", s.getArea())
}

func main() {
	// UP CASTING
	var shape1 Shape = Square{side: 5}
	var shape2 Shape = Rectangle{width: 3, length: 4}
	var shape3 Shape = Rectangle{width: 7, length: 9}

	fmt.Printf("shape1: %#v \n", shape1)
	fmt.Printf("shape2: %#v \n", shape2)
	fmt.Printf("shape3: %#v \n", shape3)

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes {
		fmt.Printf("%#v memiliki luas %.1f dan keliling %.1f \n", shape, shape.getArea(), shape.getPerimeter())
	}

	getAreaOfRectangle(Rectangle{width: 3.2, length: 5.1})
	getAreaOfSquare(Square{side: 4.3})
	getArea(Rectangle{width: 1.1, length: 1.2})
	getArea(Square{side: 1.5})

	var square1 Shape = Square{side: 6.7}
	fmt.Println("getArea:", square1.getArea())
	fmt.Println("getPerimeter:", square1.getPerimeter())
	// fmt.Println("side:",square1.side) TIDAK BISA DILAKUKAN

	// DOWN CASTING
	var originalSquare Square = square1.(Square)
	fmt.Println("side:", originalSquare.side)
	fmt.Println("getArea:", originalSquare.getArea())

	// Interface kosong
	// var anything interface{}
	var anything any
	anything = "Faisa"
	fmt.Printf("anything %T: %v \n", anything, anything)
	anything = 23
	fmt.Printf("anything %T: %v \n", anything, anything)
	anything = false
	fmt.Printf("anything %T: %v \n", anything, anything)
	anything = []string{"Golang", "Java", "Python"}
	fmt.Printf("anything %T: %v \n", anything, anything)

}
