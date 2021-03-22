package testing

/*
 * This file will add some types to the database.
 */

import (
	"github.com/cagox/fluxspellsapi/app/models"
)

func addTypes() []int {
	types := [...]models.Type{
		{Name: "Archery", Description: "Archery Spells", Summary: "Archery Spells"},
		{Name: "Cooking", Description: "Cooking Spells", Summary: "Cooking Spells"},
		{Name: "Information", Description: "Spells that give you information", Summary: "Information Spells"},
		{Name: "Area Of Effect", Description: "Spells that effect more than one target", Summary: "AoE Spells"},
		{Name: "Combat", Description: "Spells for Attacking", Summary: "Attack Spells"},
		{Name: "Defense", Description: "Spells for Defending", Summary: "Defense Spells"},
		{Name: "Utility", Description: "Utility Spells", Summary: "Utility Spells"},
	}

	indexes := make([]int, 0)

	for _, value := range types {
		newValue := models.InsertType(&value)
		indexes = append(indexes, newValue.ID)
	}

	return indexes

}
