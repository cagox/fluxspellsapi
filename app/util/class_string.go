package util

/*
 * ClassString returns an empty string if classNames is empty.
 * Otherwise it returns a string in the form of `class="one or more classes"`
 * This method is meant to facilitate generating HTML code.
 */
func ClassString(classNames ...string) string {
	classes := ""
	if len(classNames) != 0 {
		classes += `class="`
		for i, name := range classNames {
			if i == 0 {
				classes += name
			} else {
				classes += " " + name
			}
		}
		classes += `"`
	}
	return classes
}

//TODO: Look at this again some time from a proper IDE. It's simple enough that it should be fine though.
