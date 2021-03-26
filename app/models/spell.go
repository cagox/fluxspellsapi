package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/util"
)

/*
 * Spell is the full spell struct that will be used to display spell information.
 */
type Spell struct {
	ID            int      `json:"spell_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Summary       string   `json:"summary"`
	Cost          string   `json:"cost"`
	Difficulty    string   `json:"difficulty"`
	Prerequisites string   `json:"prerequisites"`
	Components    string   `json:"components"`
	Range         string   `json:"spellrange"`
	Types         []Type   `json:"types"`
	Schools       []School `json:"schools"`
	/*
	 * Types and Schools are not persisted in the spells table in the database.
	 * They are collected using the GetSchools and GetTypes methods below.
	 */
}

/*
 * SpellSummary is the struct that will be used to display brief information on index pages.
 */
type SpellSummary struct {
	ID      int    `json:"spell_id"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

func (spell Spell) Link(classNames ...string) string {
	classes := util.ClassString(classNames...)
	return fmt.Sprintf(`<a %s href="/spells/%d">%s</a>`, classes, spell.ID, spell.Name)
}

func (spell SpellSummary) Link(classNames ...string) string {
	classes := util.ClassString(classNames...)
	return fmt.Sprintf(`<a %s href="/spells/%d">%s</a>`, classes, spell.ID, spell.Name)
}

func CreateAndInsertSpell(name string, description string, summary string, cost string, difficulty string, prerequisites string, components string, spellRange string) *Spell {
	spell := &Spell{Name: name, Description: description, Summary: summary, Cost: cost, Difficulty: difficulty, Prerequisites: prerequisites, Components: components, Range: spellRange}
	return InsertSpell(spell)
	//TODO: Add error checking.
}

func InsertSpell(spell *Spell) *Spell {
	sqlString := `INSERT INTO spells(name, description, summary,cost, difficulty, prerequisites, components, spellrange) VALUES(?,?,?,?,?,?,?,?)`

	res, err := app.Config.Database.Exec(sqlString, spell.Name, spell.Description, spell.Summary, spell.Cost, spell.Difficulty, spell.Prerequisites, spell.Components, spell.Range)
	if err != nil {
		panic(err) //TODO: More effective error checking.
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //TODO: More effective error checking.
	}

	spell.ID = int(id)
	return spell
}

func (spell Spell) GetTypes() []Type {
	spellTypes := make([]Type, 0)

	sqlStatement := `SELECT types.type_id, types.name, types.summary FROM types INNER JOIN type_links ON types.type_id = type_links.type_id WHERE type_links.spell_id=?`

	rows, err := app.Config.Database.Query(sqlStatement, spell.ID)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextType := Type{}
		if err := rows.Scan(&nextType.ID, &nextType.Name, &nextType.Summary); err != nil {
			panic(err) //TODO: Figure out how to handle this properly
		}
		spellTypes = append(spellTypes, nextType)
	}
	return spellTypes
}

func (spell Spell) GetSchools() []School {
	schools := make([]School, 0)

	sqlStatement := `SELECT schools.school_id, schools.name, schools.summary FROM schools INNER JOIN school_links ON schools.school_id = school_links.school_id WHERE school_links.spell_id=?`

	rows, err := app.Config.Database.Query(sqlStatement, spell.ID)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextSchool := School{}
		if err := rows.Scan(&nextSchool.ID, &nextSchool.Name, &nextSchool.Summary); err != nil {
			panic(err) //TODO: Figure out how to handle this properly
		}
		schools = append(schools, nextSchool)
	}
	return schools
}

func (spell Spell) AddType(typeID int) (*TypeLink, error) {
	return InsertTypeLink(spell.ID, typeID)
}

func (spell Spell) AddSchool(schoolID int) (*SchoolLink, error) {
	return InsertSchoolLink(spell.ID, schoolID)
}

func GetSpellByID(id int) *Spell {
	spell := new(Spell)

	sqlStatement := `SELECT spell_id, name, description, summary, cost, difficulty, prerequisites, components, spellrange FROM spells WHERE spell_id=?`

	row := app.Config.Database.QueryRow(sqlStatement, id)
	switch err := row.Scan(&spell.ID, &spell.Name, &spell.Description, &spell.Summary, &spell.Cost, &spell.Difficulty, &spell.Prerequisites, &spell.Components, &spell.Range); err {
	case sql.ErrNoRows:
		fmt.Println("Spell ID ", id, "doesn't exist.")
	case nil:
		return spell
	default:
		fmt.Println(`GetSpellByID(id int)`)
		panic(err)
	}
	return spell
	//TODO: Fix error checking to something more useful for web dev.
}

func (spell *Spell) ToJSON() string {
	spell.Schools = spell.GetSchools()
	spell.Types = spell.GetTypes()

	spellJSON, err := json.Marshal(spell)
	if err != nil {
		panic(err)
	}
	return string(spellJSON)
}

func (spell *Spell) TypesToJSON() string {
	types := spell.GetTypes()
	typesList, err := json.Marshal(types)
	if err != nil {
		panic(err)
	}
	return string(typesList)
}

func (spell *Spell) SchoolsToJSON() string {
	schools := spell.GetSchools()
	schoolsList, err := json.Marshal(schools)
	if err != nil {
		panic(err)
	}
	return string(schoolsList)
}
