package controller

import "fmt"

func HelloWorld(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
