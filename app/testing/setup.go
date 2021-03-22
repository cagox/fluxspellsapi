package testing

/*
 * This will use the data from the other files to get things up and going.
 */

func Setup() {
	schoolIDs := addSchools()
	typeIDs := addTypes()
	spellIDs := addSpells()
	linkSchools(spellIDs, schoolIDs)
	linkTypes(spellIDs, typeIDs)
}
