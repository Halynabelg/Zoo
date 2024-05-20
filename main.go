package main

import "fmt"

type Zookeeper struct {
	Animal     uint
	Cage       uint
	Title      string
	Desciption string
}

func (s Zookeeper) Print() {
	fmt.Printf("%s(%d-%d)\n%s\n", s.Title, s.Animal, s.Cage, s.Desciption)

}
func main() {
	s := Zookeeper{
		Animal:     8,
		Cage:       5,
		Title:      "Escaped from the Zoo animals from cages ",
		Desciption: "Zookeeper must catch all the animals until the morning",
	}
	s.Print()
}
