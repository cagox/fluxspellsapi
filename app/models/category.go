package models

import (
	"database/sql"
	//"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
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

func GetCategoryList() []CategorySummary {
	categories := make([]CategorySummary, 0)

	sqlStatement := `SELECT category_id, name, summary FROM categories;`
	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		fmt.Println(`Failed at GetCategoryList(): Database.Query`)
		panic(err) //TODO: Make this more useful.
	}
	defer rows.Close()

	for rows.Next() {
		nextCategory := CategorySummary{}

		if err := rows.Scan(&nextCategory.CategoryID, &nextCategory.Name, &nextCategory.Summary); err != nil {
			fmt.Println(`Failed at GetCategoryList(): rows.Scan`)
			panic(err) //TODO: Make this more useful.
		}

		categories = append(categories, nextCategory)
	}
	return categories

}

func GetCategoryListAsJSON() []byte {
	categories := GetCategoryList()

	categoryList, err := json.Marshal(categories)
	if err != nil {
		fmt.Println("GetCategoryListAsJSON: Failed to marshal schoolList")
	}

	return categoryList
}

func GetCategoryByID(id int) *Category {
	category := new(Category)

	sqlStatement := `SELECT category_id, name, summary, description FROM categories WHERE category_id =?`

	row := app.Config.Database.QueryRow(sqlStatement, id)

	switch err := row.Scan(&category.CategoryID, &category.Name, &category.Summary, &category.Description); err {
	case sql.ErrNoRows:
		fmt.Println("Category ID ", id, " doesn't exist.")
	case nil:
		return category
	default:
		fmt.Println(`GetCategoryByID(id int)`)
		panic(err)
	}
	return category
	//TODO: Fix error checking to do something more useful.
}

func GetCategoryAsJSON(id int) []byte {
	category := GetCategoryByID(id)

	categoryJSON, err := json.Marshal(category)

	if err != nil {
		panic(err) //TODO: Make this more useful
	}

	return categoryJSON
}
