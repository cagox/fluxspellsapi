package models

import (
	"encoding/gob"
	"encoding/json"
	"github.com/cagox/fluxspellsapi/app"
)

type Spell struct {
	SpellID        int        `json:"spell_id"`
	Name           string     `json:"name"`
	Cost           string     `json:"cost"`
	Difficulty     string     `json:"difficulty"`
	SpellRange     string     `json:"spellrange"`
	Prerequisites  string     `json:"prerequisites"`
	AbilityScoreID int        `json:"ability_score_id"`
	Summary        string     `json:"summary"`
	Description    string     `json:"description"`
	Categories     []Category `json:"categories"`
	Schools        []School   `json:"schools"`
	/*
	 * Categories and Schools are not persisted in the spells table in the database.
	 * They are collected using the GetSchools and GetCategories methods below.
	 */
}

type SpellSummary struct {
	SpellID    int              `json:"spell_id"`
	Name       string           `json:"name"`
	Summary    string           `json:"summary"`
	Categories []CategoryHeader `json:"categories"`
	Schools    []SchoolHeader   `json:"schools"`
}

func init() {
	gob.Register(Spell{})
	gob.Register(SpellSummary{})
}

func GetSpellSummaryList() []SpellSummary {
	spells := make([]SpellSummary, 0)

	sqlStatement := `SELECT spell_id, name, summary FROM spells;`
	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		panic(err) //TODO: build up proper error handling.
	}
	defer rows.Close()

	for rows.Next() {
		nextSpell := SpellSummary{}
		if err := rows.Scan(&nextSpell.SpellID, &nextSpell.Name, &nextSpell.Summary); err != nil {
			panic(err)
			//TODO: Figure out proper error checking and logging.
		}
		nextSpell.Schools = nextSpell.GetSchools()
		nextSpell.Categories = nextSpell.GetCategories()
		spells = append(spells, nextSpell)
	}

	if err != nil {
		panic(err)
		//TODO: Figure out proper error checking and logging.
	}

	return spells
}

func GetSpellSummaryListAsJSON() []byte {
	spells := GetSpellSummaryList()

	spellList, err := json.Marshal(spells)
	if err != nil {
		panic(err)
		//TODO: Figure out proper error checking and logging.
	}

	return spellList
}

func (spell SpellSummary) GetSchools() []SchoolHeader {
	schools := make([]SchoolHeader, 0)

	sqlStatement := `SELECT schools.school_id, schools.name FROM schools INNER JOIN spellschools ON schools.school_id = spellschools.school_id WHERE spellschools.spell_id=?`

	rows, err := app.Config.Database.Query(sqlStatement, spell.SpellID)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextSchool := SchoolHeader{}
		if err := rows.Scan(&nextSchool.SchoolID, &nextSchool.Name); err != nil {
			panic(err) //TODO: Figure out how to handle this properly
		}
		schools = append(schools, nextSchool)
	}
	return schools
}

func (spell SpellSummary) GetCategories() []CategoryHeader {
	categories := make([]CategoryHeader, 0)

	sqlStatement := `SELECT categories.category_id, categories.name FROM categories INNER JOIN spellcategories ON categories.category_id = spellcategories.category_id WHERE spellcategories.spell_id=?`

	rows, err := app.Config.Database.Query(sqlStatement, spell.SpellID)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextCategory := CategoryHeader{}
		if err := rows.Scan(&nextCategory.CategoryID, &nextCategory.Name); err != nil {
			panic(err) //TODO: Figure out how to handle this properly
		}
		categories = append(categories, nextCategory)
	}
	return categories
}

//GetSchools This function builds on what we already did for the summary.
func (spell Spell) GetSchools() []SchoolHeader {
	tempSpell := SpellSummary{SpellID: spell.SpellID}
	return tempSpell.GetSchools()
}

func (spell Spell) GetCategories() []CategoryHeader {
	tempSpell := SpellSummary{SpellID: spell.SpellID}
	return tempSpell.GetCategories()
}

func GetSpellsBySchool(school_id int) []SpellSummary {
	spells := make([]SpellSummary, 0)

	sqlStatement := `SELECT spells.spell_id, spells.name, spells.summary FROM spells INNER JOIN spellschools on spells.spell_id = spellschools.spell_id WHERE spellschools.school_id=?`

	rows, err := app.Config.Database.Query(sqlStatement, school_id)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextSpell := SpellSummary{}
		if err := rows.Scan(&nextSpell.SpellID, &nextSpell.Name, &nextSpell.Summary); err != nil {
			panic(err) //TODO: Figure out how to handle this properly
		}
		nextSpell.Schools = nextSpell.GetSchools()
		nextSpell.Categories = nextSpell.GetCategories()
		spells = append(spells, nextSpell)
	}
	return spells
}

func GetSpellsBySchoolAsJSON(school_id int) []byte {
	spells := GetSpellsBySchool(school_id)

	spellList, err := json.Marshal(spells)
	if err != nil {
		panic(err)
		//TODO: Figure out proper error checking and logging.
	}

	return spellList
}

func GetSpellsByCategory(category_id int) []SpellSummary {
	spells := make([]SpellSummary, 0)

	sqlStatement := `SELECT spells.spell_id, spells.name, spells.summary FROM spells INNER JOIN spellcategories on spells.spell_id = spellcategories.spell_id WHERE spellcategories.category_id=?`

	rows, err := app.Config.Database.Query(sqlStatement, category_id)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextSpell := SpellSummary{}
		if err := rows.Scan(&nextSpell.SpellID, &nextSpell.Name, &nextSpell.Summary); err != nil {
			panic(err) //TODO: Figure out how to handle this properly
		}
		nextSpell.Schools = nextSpell.GetSchools()
		nextSpell.Categories = nextSpell.GetCategories()
		spells = append(spells, nextSpell)
	}
	return spells
}

func GetSpellsByCategoryAsJSON(category_id int) []byte {
	spells := GetSpellsByCategory(category_id)

	spellList, err := json.Marshal(spells)
	if err != nil {
		panic(err)
		//TODO: Figure out proper error checking and logging.
	}

	return spellList
}
