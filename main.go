package main

import "fmt"

func main() {
	a := 2
	b := 5
	// 변수명 앞에 & 붙이면 메모리 주소
	fmt.Println(&a, &b)

	c := 2
	d := &c
	*d = 20
	fmt.Println(c)
}
