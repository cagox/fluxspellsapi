package main

import (
	"bufio"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/database"
	"github.com/cagox/fluxspellsapi/app/models"
	"os"
	"strings"
)

func main() {
	app.StartLogging()
	defer app.StopLogging()
	database.OpenDatabase()
	defer database.CloseDatabase()

	fmt.Println("Hello World!")
	fmt.Println("Welcome to " + app.Config.SiteName)

	//login()
	//testing.Setup() //This has already been done, so I theoretically should not have to do it again.

	testSpell := models.GetSpellByID(1)

	fmt.Println(testSpell.Name + " Spell: " + testSpell.ToJSON())
	fmt.Println()

	testSchool := models.GetSchoolByID(17)

	fmt.Println("Spells for the School " + testSchool.Name + ": " + testSchool.SpellsToJSON())
	fmt.Println()

	testType := models.GetTypeByID(21)
	fmt.Println(testType.Name + " type spells:" + testType.SpellsToJSON())

	testSpell.AddType(21) //Already done, and it works.
	fmt.Println(testSpell.Name + " has the following types: " + testSpell.TypesToJSON())

}

func login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	password = strings.TrimSuffix(password, "\n")
	email = strings.TrimSuffix(email, "\n")

	user := models.GetUserByEmail(email)

	authenticated := user.AuthenticateUser(password)

	if authenticated {
		fmt.Println(user.Email + " logged in!")
	} else {
		fmt.Println("Wrong Password!")
	}

}
