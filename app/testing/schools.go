package testing

/*
 * This package will add some schools to the database.
 */

import (
	"github.com/cagox/fluxspellsapi/app/models"
)

func addSchools() []int {
	schools := [...]models.School{
		{Name: "Earth", Description: "Earth Elemental Magic", Summary: "Earth Elemental Magic"},
		{Name: "Air", Description: "Air Elemental Magic", Summary: "Air Elemental Magic"},
		{Name: "Water", Description: "Water Elemental Magic", Summary: "Water Elemental Magic"},
		{Name: "Fire", Description: "Fire Elemental Magic", Summary: "Fire Elemental Magic"},
		{Name: "Life", Description: "Life and Healing Magic", Summary: "Life and Healing Magic"},
		{Name: "Light", Description: "Light Magic", Summary: "Light Magic"},
		{Name: "Arcane", Description: "Powerful Magic that is hard to categorize.", Summary: "Arcane Magic"},
		{Name: "Food", Description: "Magic relating to food.", Summary: "Food Magic"},
	}

	indexes := make([]int, 0)

	for _, value := range schools {
		newValue := models.InsertSchool(&value)
		indexes = append(indexes, newValue.ID)
	}

	return indexes

}
