package main

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

var myName string = "Anderson Teala"
var initName string = "Me"

func personIsASalesman(name string) bool {
	person := string(name[0]) + string(name[1])
	if person == initName {
		return true
	} else {
		return false
	}
}

func generateFakeName(quantityNames int) []string {
	fakeNamesMap := []string{myName}

	gofakeit.Seed(0)
	for i := 0; i < quantityNames; i++ {
		fakeNamesMap = append(fakeNamesMap, gofakeit.Name())
	}

	return fakeNamesMap
}

func findSeller(fakeNames []string) bool {
	verifiedNames := []string{}
	for index := 0; index < len(fakeNames); index++ {
		name := fakeNames[index]
		personHasBeenVerified := strings.Contains(strings.Join(verifiedNames, ","), name)
		if !personHasBeenVerified {
			if personIsASalesman(name) {
				fmt.Println("Person: "+name+", found in the index", index)
				return true
			} else {
				verifiedNames = append(verifiedNames, name)
				namesGenerated := generateFakeName(10)
				fakeNames = append(fakeNames, namesGenerated...)
			}
		}
	}

	fmt.Println("Person not found")
	return false
}

func main() {
	fakeNames := generateFakeName(10)
	findSeller(fakeNames)
}
