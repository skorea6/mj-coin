package main

import (
	"fmt"
	"mjcoin/person"
)

func golang() {
	fmt.Println("It Works!")

	nico := person.Person{}
	nico.SetDetails("Nico", 42)
	fmt.Println("Main ", nico)
	fmt.Println(nico.Name())
}
