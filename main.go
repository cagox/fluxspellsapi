package main

import (
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(app.Config.SiteName)
}
