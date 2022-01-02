package main


import (
	"fmt"
	"apiProject/app"
)

func main() {
	//Intialise app
	app := app.App{}
	app.Initialize(
		"username",
		"password",
		"databasename")

	//run the application
	fmt.Println("starting the application...")
	app.Run("121.0.0.1:8010")

}
