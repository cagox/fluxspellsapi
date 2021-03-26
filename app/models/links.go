package models

import (
	"encoding/json"
	"github.com/VividCortex/mysqlerr"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/go-sql-driver/mysql"
)

type TypeLink struct {
	ID      int `json:"id"`
	TypeID  int `json:"type_id"`
	SpellID int `json:"spell_id"`
}

type SchoolLink struct {
	ID       int `json:"id"`
	SchoolID int `json:"school_id"`
	SpellID  int `json:"spell_id"`
}

func InsertTypeLink(spellID int, typeID int) (*TypeLink, error) {
	newLink := &TypeLink{SpellID: spellID, TypeID: typeID}

	sqlString := `INSERT INTO type_links(type_id, spell_id) VALUES(?,?)`
	res, err := app.Config.Database.Exec(sqlString, typeID, spellID)
	if err != nil {
		if theError, ok := err.(*mysql.MySQLError); ok {
			switch theError.Number {
			case mysqlerr.ER_DUP_ENTRY:
				return nil, err
			default:
				panic(err)
			}
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //TODO: Proper error checking.
	}

	newLink.ID = int(id)
	return newLink, nil
}

func InsertSchoolLink(spellID int, schoolID int) (*SchoolLink, error) {
	newLink := &SchoolLink{SpellID: spellID, SchoolID: schoolID}

	sqlString := `INSERT INTO school_links(school_id, spell_id) VALUES(?,?)`
	res, err := app.Config.Database.Exec(sqlString, schoolID, spellID)
	if err != nil {
		if theError, ok := err.(*mysql.MySQLError); ok {
			switch theError.Number {
			case mysqlerr.ER_DUP_ENTRY:
				return nil, err
			default:
				panic(err)
			}
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //TODO: Proper error checking.
	}

	newLink.ID = int(id)
	return newLink, nil
}

//TODO: Add the ability to drop links.
func (link *SchoolLink) ToJSON() string {
	linkJson, err := json.Marshal(link)
	if err != nil {
		panic(err)
	}

	return string(linkJson)
}

func (link *TypeLink) ToJSON() string {
	linkJson, err := json.Marshal(link)
	if err != nil {
		panic(err)
	}

	return string(linkJson)
}
