package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/paddycakes/arranmore-api/internal/sensor"
	transportHTTP "github.com/paddycakes/arranmore-api/internal/transport/http"
	"net/http"
)

// App - the struct which contains things
// like pointers to database connections
type App struct {}

// Run - sets up Arranmore REST API
func (app *App) Run() error  {
	fmt.Println("Setting up Arranore REST API")

	//var err error
	//_, err = database.NewDatabase()
	//if err != nil {
	//	return err
	//}

	sensorService := sensor.NewService()

	handler := transportHTTP.NewHandler(sensorService)
	handler.SetupRoutes()

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// Need to sort this out
	// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(handler.Router)); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main()  {
	fmt.Println("Arranmore REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up Arranmore REST API")
		fmt.Println(err)
	}
}


