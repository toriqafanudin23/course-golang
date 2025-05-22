package main

import "fmt"

type Celcius float64

type Patient struct {
	Name string
	Age  int
	Celcius
}

func (c Celcius) toFahrenheit() float64 {
	return float64(c*9/5 + 32)
}

func (c *Celcius) add(value float64) {
	*c += Celcius(value)
}

func main() {
	var temperature Celcius = 24.6
	fmt.Println("temperature:", temperature)
	fmt.Println(temperature, "Celcius =", temperature.toFahrenheit(), "Fahrenheit")
	temperature.add(3)
	fmt.Println("temperature now:", temperature)
	newPatient := Patient{Name: "Faisa", Age: 25, Celcius: 39.7}
	fmt.Printf("new patient: %+v \n", newPatient)
}
