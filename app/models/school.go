package models

import (
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
