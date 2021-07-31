package inventorykata

import "errors"

// Categories
type Category int

const (
	allAboutNuts Category = iota
	allAboutCandies
	nutsAndCandies
)

func (c Category) IsValid() error {
	switch c {
	case allAboutNuts, allAboutCandies, nutsAndCandies:
		return nil
	}
	return errors.New("invalid category")
}

func (c Category) String() string {
	return [...]string{"All About Nuts", "All About Candies", "Nuts And Candies"}[c]
}

/////

type Grade int

const (
	premium Grade = iota
	regular
)

/////

type Size struct {
	size       string
	gramWeight int64
}

/////

type Product struct {
	name        string
	description string
	category    Category
	grade       Grade
	size        Size
	price       float64
}

/////
