package models

import (
	//"database/sql"
	"encoding/gob"
	"encoding/json"
	"github.com/cagox/fluxspellsapi/app"
	//"fmt"
	//"time"
)

type Category struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type CategorySummary struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	Summary    string `json:"summary"`
}

type CategoryHeader struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

func init() {
	gob.Register(Category{})
	gob.Register(CategorySummary{})
	gob.Register(CategoryHeader{})
}

func GetCategoryHeaderList() []byte {
	categories := make([]CategoryHeader, 0)

	sqlStatement := `SELECT category_id, name FROM categories;`
	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		panic(err) //TODO: build up proper error handling.
	}
	defer rows.Close()

	for rows.Next() {
		nextCategory := CategoryHeader{}
		if err := rows.Scan(&nextCategory.CategoryID, &nextCategory.Name); err != nil {
			panic(err)
			//TODO: Figure out proper error checking and logging.
		}
		categories = append(categories, nextCategory)
	}

	categoryList, err := json.Marshal(categories)

	return categoryList
}
