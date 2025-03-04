package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plus2(a, b int, name string) (int, string) {
	return a + b, name
}

func plus3(a ...int) int {
	var total int
	for _, item := range a {
		total += item
	}
	return total
}

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Printf("Hello, %s!\n", p.name)
}

func main() {
	fmt.Println("It Works!")

	//name2 := "이렇게도 쓸수 있다. 위와 같음 타입추론함."
	age := 18
	age = 12
	fmt.Println(age)

	result := plus(2, 2)
	fmt.Println(result)

	result3 := plus3(2, 3, 4, 5, 6, 7)
	fmt.Println(result3)

	var name string = "nico"
	for index, letter := range name {
		fmt.Println(index, letter)  // byte
		fmt.Println(string(letter)) // string
		fmt.Printf("%b\n", letter)  // 2진수
	}

	x := 4848498445
	fmt.Printf("%b\n", x)
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBinary)

	// Array
	foods := [3]string{"a", "b", "c"} // 개수 명시 필수
	for _, food := range foods {
		fmt.Println(food)
	}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}

	// slice : 개수 length 무한히 커질수 있음
	foods_slice := []string{"a", "b", "c"}
	fmt.Printf("%v\n", foods_slice)
	foods_slice = append(foods_slice, "d")
	fmt.Printf("%v\n", foods_slice)
	fmt.Println(len(foods_slice))

	a := 2
	b := &a // a의 메모리 '주소' 저장
	a = 50
	fmt.Println(b, &a)
	fmt.Println(a)
	fmt.Println(*b) // &는 주소, *는 값

	nico := person{name: "Nico", age: 30}
	nico2 := person{"Nico2", 31}
	nico.sayHello()
	fmt.Println(nico2.age)
}
