package models

import (
	"encoding/gob"
)

/*
 * Spell is the full spell struct that will be used to display spell information.
 */
type Spell struct {
	ID            int    `json:"spell_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Summary       string `json:"summary"`
	Cost          string `json:"cost"`
	Difficulty    string `json:"difficulty"`
	Prerequisites string `json:"prerequisites"`
	Components    string `json:"components"`
	Range         string `json:"spellrange"`
	//Types()        Many-to-Many
	//Schools()      Many-to-Many
}



/*
 * SpellSummary is the struct that will be used to display brief information on index pages.
 */
type SpellSummary struct {
	ID       int     `json:"spell_id"`
	Name     string  `json:"name"`
	Summary  string  `json:"summary"`
}


func (spell Spell) Link(classNames ...string) string {
	classes := util.ClassString(classNames)
	return fmt.Sprintf(`<a %s href="/spells/%d">%s</a>`, classes, spell.ID, spell.Name)
}

func (spell SpellSummary) Link(classNames ...string) string {
	classes := util.ClassString(classNames)
	return fmt.Sprintf(`<a %s href="/spells/%d">%s</a>`, classes, spell.ID, spell.Name)
}


func CreateAndInsertSpell(name string, description string, summary string, cost string, difficulty string, prerequisities string, components string spellRange string) * Spell {
	spell := &Spell{Name: name, Description: description, Summary: summary, Cost: cost, Difficulty: difficulty, Prerequisites: prerequisites, Components: components, Range: spellRange}
	return InsertSpell(spell)
	//TODO: Add error checking.
}

func InsertSpell(spell *Spell) *Spell {
	sqlString := `INSERT INTO spells(?) VALUES(?,?,?,?,?,?,?,?)`
	fields := `name, description, summary,cost, difficulty, prerequisites, components, range`
	res, err := app.Config.Database.Exec(sqlString,fields, spell.Name, spell.Description, spell.Summary, spell.Cost, spell.Difficulty, spell.Prerequisites, spell.Components, spell.Range)
	if err !=  nil {
		panic(err) //TODO: More effective error checking.
	}
	
	id, err := res.LastInsertId()
		if err !=  nil {
		panic(err) //TODO: More effective error checking.
	}
	
	spell.ID = int(id)
	return spell
}


func (spell Spell) Types() []Type {
	spellTypes := make([]Type, 0)

	sqlStatement := `SELECT ? FROM types INNER JOIN type_links ON types.type_id = type_links.type_id WHERE type_links.spell_id=?`
	fields := `types.type_id, types.name, types.summary`
	
	rows, err := app.Config.Database.Query(sqlStatement, fields, spell.ID)
		if err != nil {
		//TODO: Figure out proper error checking and logging.
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextType := Type{}
		if err := rows.Scan(&nextType.ID, &nextType.Name, &nextType.Summary); err != nil {
			log.Fatal(err) //TODO: Figure out how to handle this properly
		}
		append(spellTypes, nextType)
	}
	return spellTypes
}

func (spell Spell) Schools() []School {
	schools := make([]School, 0)

	sqlStatement := `SELECT ? FROM schools INNER JOIN school_links ON schools.school_id = school_links.school_id WHERE school_links.spell_id=?`
	fields := `schools.school_id, schools.name, schools.summary`
	
	rows, err := app.Config.Database.Query(sqlStatement, fields, spell.ID)
		if err != nil {
		//TODO: Figure out proper error checking and logging.
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextSchool := School{}
		if err := rows.Scan(&nextSchool.ID, &nextSchool.Name, &nextSchool.Summary); err != nil {
			log.Fatal(err) //TODO: Figure out how to handle this properly
		}
		append(schools, nextSchool)
	}
	return schools
}

func (spell Spell) AddType(typeID int) {
	InsertTypeLink(spell.ID, typeID)
}


func (spell Spell) AddSchool(schoolID int) {
	InsertSchoolLink(spell.ID, schoolID)
}

func GetSpellByID(id int) *Spell {
	spell = new(Spell)
	
	sqlStatement := `SELECT ? FROM spells WHERE spell_id=?`
	fields := `spell_id, name, description, summary, cost, difficulty, prerequisites, components, spellrange`
	
	row, err := app.Config.Database.QueryRow(sqlStatement, fields, id)
	switch err := row.Scan(&spell.ID, &spell.Name, &spell.Description, &spell.Summary, &spell.Cost, &spell.Difficulty, &spell.Prerequisites, &spell.Components, &spell.Range); err {
	case sql.ErrNoRows:
		fmt.Println("Spell ID ", id,  "doesn't exist.")
	case nil:
		return spell
	default:
		panic(err)
	}
	return spell
	//TODO: Fix error checking to something more useful for web dev.	
}

func init() {
	gob.Register(Spell{})
	gob.Register(SpellSummary{})
}
