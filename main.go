package main

import "fmt"

// struct. method도 가능. go의 struct에는 생성자가 없음
type person struct {
	name    string
	age     int
	favFood []string
}

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

	// Map -> map[key type]value type
	// value 다른 type들 쓰려면 struct
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)

	// Map도 for문 사용 가능(key, value)
	for _, value := range nico {
		fmt.Println(value)
	}

	// struct
	favFood := []string{"kimchi", "remen"}
	pbk := person{name: "nico", age: 18, favFood: favFood} //value만 적어도 되고, key와 같이 적어도 됨. 다만 통일해야.
	fmt.Println(pbk)
}
