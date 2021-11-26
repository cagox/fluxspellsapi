package models

import (
	"database/sql"
	"fmt"

	//"database/sql"
	"encoding/gob"
	"encoding/json"
	"github.com/cagox/fluxspellsapi/app"
	//"fmt"
	//"time"
)

type School struct {
	SchoolID    int    `json:"school_id"`
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	// Spells      []SpellSummary  `json:"spells"`
	/*
	 * Spells will not be persisted in the database. It will be collected
	 * using GetSpellSummaryList()
	 */
}

type SchoolSummary struct {
	SchoolID int    `json:"school_id"`
	Name     string `json:"name"`
	Summary  string `json:"summary"`
}

type SchoolHeader struct {
	SchoolID int    `json:"school_id"`
	Name     string `json:"name"`
}

func init() {
	gob.Register(School{})
	gob.Register(SchoolSummary{})
	gob.Register(SchoolHeader{})
}

func GetSchoolHeaderList() []byte {
	schools := make([]SchoolHeader, 0)

	sqlStatement := `SELECT school_id, name FROM schools;`
	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		panic(err) //TODO: build up proper error handling.
	}
	defer rows.Close()

	for rows.Next() {
		nextSchool := SchoolHeader{}
		if err := rows.Scan(&nextSchool.SchoolID, &nextSchool.Name); err != nil {
			panic(err)
			//TODO: Figure out proper error checking and logging.
		}
		schools = append(schools, nextSchool)
	}

	schoolList, err := json.Marshal(schools)

	return schoolList
}

func GetSchoolByID(id int) *School {
	school := new(School)

	sqlStatement := `SELECT school_id, name, summary, description FROM schools WHERE school_id =?`

	row := app.Config.Database.QueryRow(sqlStatement, id)

	switch err := row.Scan(&school.SchoolID, &school.Name, &school.Summary, &school.Description); err {
	case sql.ErrNoRows:
		fmt.Println("School ID ", id, " doesn't exist.")
	case nil:
		return school
	default:
		fmt.Println(`GetSpellByID(id int)`)
		panic(err)
	}
	return school
	//TODO: Fix error checking to do something more useful.
}

func GetSchoolAsJSON(id int) []byte {
	school := GetSchoolByID(id)

	schoolJSON, err := json.Marshal(school)

	if err != nil {
		panic(err) //TODO: Make this more useful
	}

	return schoolJSON
}

func GetSchoolList() []SchoolSummary {
	schools := make([]SchoolSummary, 0)

	sqlStatement := `SELECT school_id, name, summary FROM schools;`
	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		fmt.Println(`Failed at GetSchoolList(): Database.Query`)
		panic(err) //TODO: Make this more useful.
	}
	defer rows.Close()

	for rows.Next() {
		nextSchool := SchoolSummary{}

		if err := rows.Scan(&nextSchool.SchoolID, &nextSchool.Name, &nextSchool.Summary); err != nil {
			fmt.Println(`Failed at GetSchoolList(): rows.Scan`)
			panic(err) //TODO: Make this more useful.
		}

		schools = append(schools, nextSchool)
	}
	return schools

}

func GetSchoolListAsJSON() []byte {
	schools := GetSchoolList()

	schoolList, err := json.Marshal(schools)
	if err != nil {
		fmt.Println("GetSchoolListAsJSON: Failed to marshal schoolList")
	}

	return schoolList
}
