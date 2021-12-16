package models

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
)

type AbilityScore struct {
	AbilityScoreID int64  `json:"ability_score_id"`
	Name           string `json:"name"`
}

func init() {
	gob.Register(AbilityScore{})
}

func GetAbilityScoreByID(ability_score_id int64) *AbilityScore {
	ability := new(AbilityScore)

	sqlStatement := `SELECT ability_score_id, name FROM ability_scores WHERE ability_score_id=?`
	row := app.Config.Database.QueryRow(sqlStatement, ability_score_id)

	switch err := row.Scan(&ability.AbilityScoreID, &ability.Name); err {
	case sql.ErrNoRows:
		fmt.Println("Ability Score ID ", ability_score_id, " doesn't exist.")
	case nil:
		return ability
	default:
		fmt.Println(`GetAbilityScoreByID(id int)`)
		panic(err)
	}
	return ability
}

func AbilityScoreAsJSON(ability_score_id int64) []byte {
	ability := GetAbilityScoreByID(ability_score_id)

	abilityJSON, err := json.Marshal(ability)

	if err != nil {
		panic(err) //TODO: Make this more useful
	}

	return abilityJSON

}

func GetAbilityScoreList() []AbilityScore {
	abilities := make([]AbilityScore, 0)

	sqlStatement := `SELECT ability_score_id, name FROM ability_scores;`

	rows, err := app.Config.Database.Query(sqlStatement)
	if err != nil {
		fmt.Println(`Failed at GetAbilityScoreList(): Database.Query`)
		panic(err) //TODO: Make this more useful.
	}
	defer rows.Close()

	for rows.Next() {
		nextAbility := AbilityScore{}
		if err := rows.Scan(&nextAbility.AbilityScoreID, &nextAbility.Name); err != nil {
			fmt.Println(`Failed at GetAbilityScoreList(): rows.Scan`)
			panic(err) //TODO: Make this more useful.
		}
		abilities = append(abilities, nextAbility)
	}
	return abilities
}

func AbilityScoreListAsJSON() []byte {
	abilities := GetAbilityScoreList()

	abilityList, err := json.Marshal(abilities)
	if err != nil {
		fmt.Println("Failed to get JSON list of ability scores.")
		panic(err) //TODO: Make this better.
	}

	return abilityList
}
