ppackage main

import "fmt"

type Cage struct {
	animal         string
	escaptedAnimal uint
}
type Zoo struct {
	Title      string
	Desciption string
}

func putAnimal() {
	c1 := Cage{
		animal:         "Tiger",
		escaptedAnimal: 1,
	}
	c2 := Cage{
		animal:         "Lion",
		escaptedAnimal: 2,
	}
	c3 := Cage{
		animal:         "Snake",
		escaptedAnimal: 3,
	}
	c4 := Cage{
		animal:         "Papagei",
		escaptedAnimal: 4,
	}
	fmt.Println(c1.animal, c1.escaptedAnimal, c2.animal, c2.escaptedAnimal)
	fmt.Println(c3.animal, c3.escaptedAnimal, c4.animal, c4.escaptedAnimal)
}

func main() {
	z := Zoo{
		Title:      "Escaped from the Zoo animals from cages",
		Desciption: "Zookeeper must catch all the animals until the morning",
	}

	fmt.Printf(z.Title, z.Desciption)
	putAnimal()

}
