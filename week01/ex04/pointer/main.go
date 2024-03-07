package main

import "fmt"

func main(){
i,j := 42, 2701

p := &i //i는 피연산자이고, 연산자&를 통해 i값을 참조하는 포인터 p 생성
fmt.Println(*p) //연산자*를 통해 i값 출력
*p = 21 //포인터 p를 통해 i값 변경 (p가 가리키는 주소에 있는 값이 변경)
fmt.Println(i) //result: 21