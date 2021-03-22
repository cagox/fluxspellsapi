package testing

/*
 * This file will add some spells to the database.
 */

import (
	"github.com/cagox/fluxspellsapi/app/models"
)

/*
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Summary       string   `json:"summary"`
	Cost          string   `json:"cost"`
	Difficulty    string   `json:"difficulty"`
	Prerequisites string   `json:"prerequisites"`
	Components    string   `json:"components"`
	Range         string   `json:"spellrange"`
*/

func addSpells() []int {
	spells := [...]models.Spell{
		{Name: "Flame Arrow", Description: "Launches a fiery Arrow at the enemy", Summary: "Fire Arrow", Cost: "2EP or 2SP", Difficulty: "Ranged Attack Roll", Prerequisites: "", Components: "Bow and Arrow: Arrow is consumed", Range: "Near: Spell Rank * 10 meters"},
		{Name: "Stone Skin", Description: "Turn skin to stone to reduce damage", Summary: "Stone Skin", Cost: "2EP then 1EP/round", Difficulty: "15", Prerequisites: "", Components: "", Range: "Personal"},
		{Name: "Summon Food and Drink", Description: "Summons food and water", Summary: "Summons food and water", Cost: "3EP", Difficulty: "5", Prerequisites: "", Components: "", Range: "Touch"},
		{Name: "Identify", Description: "Identify a unknown item", Summary: "Identify a unknown item", Cost: "15EP", Difficulty: "Varies", Prerequisites: "Mage Gift 1", Components: "", Range: "1 meter, +5 if touching item"},
		{Name: "Healing Circle", Description: "Heal a group of people", Summary: "Heal a group of people", Cost: "10EP", Difficulty: "20", Prerequisites: "Mage Gift 1", Components: "", Range: "POW Meters Area centered on Caster"},
	}

	indexes := make([]int, 0)

	for _, value := range spells {
		newValue := models.InsertSpell(&value)
		indexes = append(indexes, newValue.ID)
	}

	return indexes

}
