package petkata

import (
	"reflect"
	"testing"
)

func Test_getFirstNamesOfAllPeople(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"All first names", []string{"Mary", "Bob", "Ted", "Jake", "Barry", "Terry", "Harry", "John"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFns := getFirstNamesOfAllPeople(); !reflect.DeepEqual(gotFns, tt.want) {
				t.Errorf("getFirstNamesOfAllPeople() = %v, want %v", gotFns, tt.want)
			}
		})
	}
}

func Test_getNamesOfMarySmithsPets(t *testing.T) {
	tests := []struct {
		name string
		want []Pet
	}{
		{"All Mary Smith's pets", []Pet{{Cat, "Tabby", 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFns := getNamesOfMarySmithsPets(); !reflect.DeepEqual(gotFns, tt.want) {
				t.Errorf("getNamesOfMarySmithsPets() = %v, want %v", gotFns, tt.want)
			}
		})
	}
}

func Test_getNumberOfPeopleWithCats(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Number of People with cats", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFns := getNumberOfPeopleWithCats(); !reflect.DeepEqual(gotFns, tt.want) {
				t.Errorf("getNumberOfPeopleWithCats() = %v, want %v", gotFns, tt.want)
			}
		})
	}
}

func Test_getNumberOfDogs(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Number of dogs", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFns := getNumberOfDogs(); !reflect.DeepEqual(gotFns, tt.want) {
				t.Errorf("getNumberOfDogs() = %v, want %v", gotFns, tt.want)
			}
		})
	}
}
