package models

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/util"
)

type School struct {
	ID          int    `json:"school_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	//Spells    mapped to func (school School) Spells() []SpellSummary {}
}

func GetSchoolByID(id int) *School {
	school := new(School)

	sqlStatement := "SELECT school_id, name, description, summary FROM schools where school_id=?"

	row := app.Config.Database.QueryRow(sqlStatement, id)

	switch err := row.Scan(&school.ID, &school.Name, &school.Description, &school.Summary); err {
	case sql.ErrNoRows:
		fmt.Println("School ", id, " does not exist.")
	case nil:
		return school
	default:
		panic(err)

	}
	return school

	//TODO: Implement error checking so that methods consuming the results can react appropriately if the school doesn't exist.
}

func GetSchoolByName(name string) *School {
	school := new(School)

	sqlStatement := "SELECT school_id, name, description, summary FROM schools where name=?"

	row := app.Config.Database.QueryRow(sqlStatement, name)

	switch err := row.Scan(&school.ID, &school.Name, &school.Description, &school.Summary); err {
	case sql.ErrNoRows:
		fmt.Println("School " + name + " does not exist.")
	case nil:
		return school
	default:
		panic(err)

	}
	return school

	//TODO: Implement error checking so that methods consuming the results can react appropriately if the school doesn't exist.
}

func (school School) Spells() []SpellSummary {

	spells := make([]SpellSummary, 0)

	// First we create the query statement.
	sqlStatement := `SELECT spells.spell_id, spells.name, spells.summary FROM spells INNER JOIN school_links ON spells.spell_id = school_links.spell_id WHERE school_links.school_id=?`

	//We execute the query.
	rows, err := app.Config.Database.Query(sqlStatement, school.ID)
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
	//TODO: Implement error checking so that methods consuming the results can react appropriately if the school is empty.
}

func (school School) Link(classNames ...string) string {
	classes := util.ClassString(classNames...)
	return fmt.Sprintf(`<a %s href="/schools/%d">%s</a>`, classes, school.ID, school.Name)
}

func CreateAndInsertSchool(name string, description string, summary string) *School {
	school := &School{Name: name, Description: description, Summary: summary}
	return InsertSchool(school)

	//TODO: Add error checking.
}

func InsertSchool(school *School) *School {
	sqlString := `INSERT INTO schools(name, description, summary) VALUES(?,?,?)`
	res, err := app.Config.Database.Exec(sqlString, school.Name, school.Description, school.Summary)
	if err != nil {
		panic(err) //Do something else with this later.
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //Do something else with this later.
	}

	school.ID = int(id)

	return school
	//TODO: Add proper error checking.
}

func init() {
	gob.Register(School{})
}

func (school *School) SpellsToJSON() []byte {
	spells := school.Spells()
	spellsJSON, err := json.Marshal(spells)
	if err != nil {
		panic(err)
	}
	return spellsJSON
}

func GetSchools() []School {
	schools := make([]School, 0)

	sqlStatement := `SELECT school_id, name, description, summary FROM schools;`

	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		panic(err) //TODO: build up proper error handling.
	}
	defer rows.Close()

	for rows.Next() {
		nextSchool := School{}
		if err := rows.Scan(&nextSchool.ID, &nextSchool.Name, &nextSchool.Description, &nextSchool.Summary); err != nil {
			panic(err)
			//TODO: Figure out proper error checking and logging.
		}
		schools = append(schools, nextSchool)
	}

	return schools
}

func SchoolsToJSON() []byte {
	schools := GetSchools()
	schoolsList, err := json.Marshal(schools)
	if err != nil {
		panic(err)
	}
	return schoolsList
}

func (school *School) toJSON() []byte {
	thisSchool, err := json.Marshal()
	if err != nil {
		panic(err)
	}
	return thisSchool
}




