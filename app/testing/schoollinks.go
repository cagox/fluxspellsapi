package testing

/*
 * This fill will connect some schools to some spells.
 */

import (
	"github.com/cagox/fluxspellsapi/app/models"
)

func linkSchools(spells []int, schools []int) {

	models.InsertSchoolLink(spells[0], schools[3]) //Flame Arrow, Fire School
	models.InsertSchoolLink(spells[1], schools[0]) //Stone Skin, Earth School
	models.InsertSchoolLink(spells[2], schools[2]) //Summon Food and Water, Water School
	models.InsertSchoolLink(spells[2], schools[7]) //Summon Food and Water, Food School
	models.InsertSchoolLink(spells[3], schools[6]) //Identify, Arcane School
	models.InsertSchoolLink(spells[4], schools[4]) //Healing Circle, Life School

}
