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

/*********/

// type Pet struct {
// 	petType PetType
// 	name    string
// 	age     int64
// }
