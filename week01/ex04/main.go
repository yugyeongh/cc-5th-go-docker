package main

import (
	"fmt"
)

type toyota struct {
	name     string
	price    float64
	discount float64
	color    string
}

func (t *toyota) getName() string {
	return t.name
}

func (t *toyota) getPrice() float64 {
	return t.price
}

func (t *toyota) getDiscount() float64 {
	return t.discount
}

func (t *toyota) getColor() string {
	return t.color
}

func (t *toyota) setColor(color string) {
	t.color = color
}

func newCar(name string, price float64, discount float64, color string) *toyota {
	return &toyota{
		name:     name,
		price:    price,
		discount: discount,
		color:    color,
	}
}

func main() {
	car := &toyota{
		name:     "car1",
		price:    4000,
		discount: 0.9,
		color:    "white",
	}

	fmt.Println(car.getName())
	car.setColor("blue")
	fmt.Println(car.getColor())
	car.setColor("red")
	fmt.Println(car.getColor())
}
