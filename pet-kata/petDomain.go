package petkata

import (
	"errors"
)

// type Person struct {
// 	firstName string
// 	lastName  string
// 	pets      []Pet
// }

type PetType string

const (
	Cat     PetType = "CAT"
	Dog             = "DOG"
	Hamster         = "HAMSTER"
	Turtle          = "TURTLE"
	Bird            = "BIRD"
	Snake           = "SNAKE"
)

func (pt PetType) IsValid() error {
	switch pt {
	case Cat, Dog, Hamster, Turtle, Bird, Snake:
		return nil
	}
	return errors.New("invalid Pet Type")
}

// func (pt PetType) String() string {
// 	return []string{"Cat", "Dog", "Hamster", "Turtle", "Bird", "Snake"}
// }

/*********/

// type Pet struct {
// 	petType string
// 	name    string
// 	age     int64
// }

// func (p Pet) getPets {

// }
