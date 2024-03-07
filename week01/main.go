package main

import (
	"fmt"

	"github.com/yugyeongh/cc-5th-go-docker/controller"
)

func main() {
	fmt.Println("hi")

	a := 1
	if a == 1 {
		fmt.Println("a==1")
	}

	fmt.Println(controller.HelloWorld("yugyeong"))

}
