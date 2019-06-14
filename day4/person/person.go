package person

import (
	"fmt"
)

type Person struct {
	Age int
}

func New(Age int) Person {
	p := Person{Age}
	return p
}

func (p Person) PersonAge() {
	fmt.Println(p.Age)
}
