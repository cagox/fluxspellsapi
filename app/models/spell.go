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
	//Types        Many-to-Many will be mapped as a function later.
	//Schools      Many-to-Many will be mapped as a function later.
}



func Types() {
	//this is a stub. It will eventually return a list of types/categories associated with the spell.
}

func Schools() {
	//This is a stub. It will eventually return a list of schools that the spell belongs in.
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


func CreateAndInsertSpell(name string, description string, summary string, cost string, difficulty string, prerequisities string, components string spellRange string) {}

func InsertSpell(spell *Spell) {}





func init() {
	gob.Register(Spell{})
	gob.Register(SpellSummary{})
}
