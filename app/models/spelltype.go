package models

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/util"
)

type Type struct {
	ID          int    `json:"type_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	//Spells    mapped to func (spellType Type) Spells() []SpellSummary {}
}

func GetTypeByID(id int) *Type {
	spellType := new(Type)

	sqlStatement := `SELECT type_id, name, description, summary FROM types where type_id=?`

	row := app.Config.Database.QueryRow(sqlStatement, id)

	switch err := row.Scan(&spellType.ID, &spellType.Name, &spellType.Description, &spellType.Summary); err {
	case sql.ErrNoRows:
		fmt.Println("Type ", id, " does not exist.")
	case nil:
		return spellType
	default:
		panic(err)

	}
	return spellType

	//TODO: Implement error checking so that methods consuming the results can react appropriately if the type doesn't exist.
}

func GetTypeByName(name string) *Type {
	spellType := new(Type)

	sqlStatement := `SELECT type_id, name, description, summary FROM types where name=?`

	row := app.Config.Database.QueryRow(sqlStatement, name)

	switch err := row.Scan(&spellType.ID, &spellType.Name, &spellType.Description, &spellType.Summary); err {
	case sql.ErrNoRows:
		fmt.Println("Type " + name + " does not exist.")
	case nil:
		return spellType
	default:
		panic(err)

	}
	return spellType

	//TODO: Implement error checking so that methods consuming the results can react appropriately if the type doesn't exist.
}

func (spellType Type) Spells() []SpellSummary {

	spells := make([]SpellSummary, 0)

	// First we create the query statement.
	sqlStatement := `SELECT spells.spell_id, spells.name, spells.summary FROM spells INNER JOIN type_links ON spells.spell_id = type_links.spell_id WHERE type_links.type_id=?`

	//We execute the query.
	rows, err := app.Config.Database.Query(sqlStatement, spellType.ID)
	if err != nil {
		//TODO: Figure out proper error checking and logging.
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		nextSpell := SpellSummary{}
		if err := rows.Scan(&nextSpell.ID, &nextSpell.Name, &nextSpell.Summary); err != nil {
			panic(err)
			//TODO: Figure out proper error checking and logging.
		}
		spells = append(spells, nextSpell)
	}

	return spells
	//TODO: Implement error checking so that methods consuming the results can react appropriately if the type is empty.
}

func (spellType Type) Link(classNames ...string) string {
	classes := util.ClassString(classNames...)
	return fmt.Sprintf(`<a %s href="/types/%d">%s</a>`, classes, spellType.ID, spellType.Name)
}

func CreateAndInsertType(name string, description string, summary string) *Type {
	spellType := &Type{Name: name, Description: description, Summary: summary}
	return InsertType(spellType)

	//TODO: Add error checking.
}

func InsertType(spellType *Type) *Type {
	sqlString := `INSERT INTO types(name, description, summary) VALUES(?,?,?)`
	res, err := app.Config.Database.Exec(sqlString, spellType.Name, spellType.Description, spellType.Summary)
	if err != nil {
		panic(err) //Do something else with this later.
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //Do something else with this later.
	}

	spellType.ID = int(id)

	return spellType
	//TODO: Add proper error checking.
}

func init() {
	gob.Register(Type{})
}

func (spellType *Type) SpellsToJSON() string {
	spells := spellType.Spells()
	spellsJSON, err := json.Marshal(spells)
	if err != nil {
		panic(err)
	}
	return string(spellsJSON)
}
