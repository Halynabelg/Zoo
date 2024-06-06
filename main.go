package main

import "fmt"

type Animal struct {
	name string
	size string
}
type Cage struct {
	Animal
	escaptedAnimal uint
}
type Zoo struct {
	Title      string
	Desciption string
}

func putAnimal() {
	c1 := Cage{
		Animal: Animal{
			name: "Tiger",
			size: "grand",
		},
		escaptedAnimal: 1,
	}
	c2 := Cage{
		Animal: Animal{
			name: "Lion",
			size: "grand",
		},
		escaptedAnimal: 2,
	}
	c3 := Cage{
		Animal: Animal{
			name: "Snake",
			size: "small",
		},
		escaptedAnimal: 3,
	}
	c4 := Cage{
		Animal: Animal{
			name: "Papagei",
			size: "small",
		},
		escaptedAnimal: 4,
	}
	fmt.Println(c1, c2, c2, c3, c4)
}

func main() {
	z := Zoo{
		Title:      "Escaped from the Zoo animals from cages",
		Desciption: "Zookeeper must catch all the animals until the morning",
	}

	fmt.Printf(z.Title)
	putAnimal()
	fmt.Printf(z.Desciption)
}
