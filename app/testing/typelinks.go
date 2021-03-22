package testing

import "github.com/cagox/fluxspellsapi/app/models"

/*
 * This fill will connect some types to some spells.
 */

func linkTypes(spells []int, types []int) {
	models.InsertTypeLink(spells[0], types[0]) //Flame Arrow, Archery Type
	models.InsertTypeLink(spells[0], types[4]) //Flame Arrow, Combat Type
	models.InsertTypeLink(spells[1], types[5]) //Stone Skin, Defense Type
	models.InsertTypeLink(spells[2], types[1]) //Summon Food and Water, Cooking Type
	models.InsertTypeLink(spells[2], types[6]) //Summon Food and Water, Utility Type
	models.InsertTypeLink(spells[3], types[2]) //Identify, Information Type
	models.InsertTypeLink(spells[3], types[6]) //Identify, Utility Type
	models.InsertTypeLink(spells[4], types[3]) //Healing Circle, AoE Type
}
