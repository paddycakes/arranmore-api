package main

import "fmt"

// App - the struct which contains things
// like pointers to database connections
type App struct {}

// Run - sets up Arranmore REST API
func (app *App) Run() error  {
	fmt.Println("Setting up Arranore REST API")
	return nil
}

func main()  {
	fmt.Println("Arranmore REST API")
	app := App{}
	if err := app.Run; err != nil {
		fmt.Println("Error starting up Arranmore REST API")
		fmt.Println(err)
	}
}


