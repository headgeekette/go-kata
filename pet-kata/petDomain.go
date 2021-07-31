package petkata

import (
	"errors"
)

// type Person struct {
// 	firstName string
// 	lastName  string
// 	pets      []Pet
// }

type PetType int

const (
	Cat PetType = iota
	Dog
	Hamster
	Turtle
	Bird
	Snake
)

func (pt PetType) IsValid() error {
	switch pt {
	case Cat, Dog, Hamster, Turtle, Bird, Snake:
		return nil
	}
	return errors.New("invalid Pet Type")
}

func (pt PetType) String() string {
	return [...]string{"Cat", "Dog", "Hamster", "Turtle", "Bird", "Snake"}[pt]
}

type Pet struct {
	petType PetType
	name    string
	age     int64
}

type Person struct {
	firstName string
	lastName  string
	Pets      []Pet
}

func getCustomerList() []Person {
	cl := make([]Person, 8)
	cl[0] = Person{"Mary", "Smith", []Pet{
		{Cat, "Tabby", 2},
	}}
	cl[1] = Person{"Bob", "Smith", []Pet{
		{Cat, "Dolly", 3},
		{Dog, "Spot", 2},
	}}
	cl[2] = Person{"Ted", "Smith", []Pet{
		{Dog, "Spike", 4},
	}}
	cl[3] = Person{"Jake", "Snake", []Pet{
		{Snake, "Serpy", 1},
	}}
	cl[4] = Person{"Barry", "Bird", []Pet{
		{Bird, "Tweety", 2},
	}}
	cl[5] = Person{"Terry", "Turtle", []Pet{
		{Turtle, "Speedy", 1},
	}}
	cl[6] = Person{"Harry", "Hamster", []Pet{
		{Hamster, "Fuzzy", 1},
		{Hamster, "Wuzzy", 1},
	}}
	cl[7] = Person{"John", "Doe", []Pet{}}

	return cl
}

func getPersonNamed(fn, ln string) Person {
	var p Person
	for _, v := range getCustomerList() {
		if v.firstName == fn && v.lastName == ln {
			return v
		}
	}
	return p
}
