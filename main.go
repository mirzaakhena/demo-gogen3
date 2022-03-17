package main

import (
	"flag"
	"fmt"
	"your/path/project/application"
	"your/path/project/shared/driver"
)

func main() {
	appMap := map[string]func() driver.RegistryContract{
		"myperson": application.NewMyperson(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
