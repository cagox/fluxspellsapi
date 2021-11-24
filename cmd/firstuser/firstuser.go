package main

import (
	"bufio"
	"fmt"
	"github.com/cagox/fluxspellsapi/app/database"
	"github.com/cagox/fluxspellsapi/app/models"
	"os"
	"strings"
)

func main() {
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

	database.OpenDatabase()
	defer database.CloseDatabase()

	firstUser := models.CreateAndInsertUser(email, password, true, true)

	fmt.Println("User "+firstUser.Email+" created with ID ", firstUser.User_ID)

}
