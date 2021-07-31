package petkata

func getFirstNamesOfAllPeople() (fns []string) {
	cl := getCustomerList()
	for _, v := range cl {
		fns = append(fns, v.firstName)
	}
	return fns
}

func getNamesOfMarySmithsPets() []Pet {
	cl := getCustomerList()
	for _, v := range cl {
		if v.firstName == "Mary" && v.lastName == "Smith" {
			return v.Pets
		}
	}
	return nil
}

func getNumberOfPeopleWithCats() int64 {
	var cp int64
	var pt PetType = Cat
	cl := getCustomerList()
	for _, v1 := range cl {
		for _, v2 := range v1.Pets {
			if v2.petType == pt {
				cp++
				break
			}
		}
	}
	return cp
}

func getNumberOfDogs() int64 {
	var d int64
	var pt PetType = Dog
	cl := getCustomerList()
	for _, v1 := range cl {
		for _, v2 := range v1.Pets {
			if v2.petType == pt {
				d++
			}
		}
	}
	return d
}
