package main

import (
	"fmt"

	"github.com/drinkwale/nomad-coin/person"
)

func main() {
	jack := person.Person{}
	jack.SetDetails("jack", 34)
	fmt.Println("Main's jack : ", jack)
}
