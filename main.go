package main

import "fmt"

func main() {
	// & : address(주소), * : 주소에 담긴 값 확인. 주소(&)에 * 붙여서 담긴 값 수정 가능
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)

	// 배열 생성(array) -> 가변 생성시 그냥 [] (slice)
	names := []string{"nico", "lynn", "dal"}
	// 배열 요소 추가 append
	names = append(names, "flynn")
	fmt.Println(names)
}
