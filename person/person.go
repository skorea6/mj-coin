package person

import "fmt"

type Person struct {
	name string
	age  int
}

// * 안붙이면, 그저 struct을 copy하는 것. 실제 main에 위치한 nico 메모리를 가르키지 않음
// *를 붙이면, 실제 main에 위치한 nico 메모리 주소를 가리킴 (수정이 가능)
func (p *Person) SetDetails(name string, age int) {
	p.age = age
	p.name = name
	fmt.Println("Setting details ", *p)
}

func (p Person) Name() string {
	return p.name
}
