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
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// Need to sort this out
	// origins := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	origins := handlers.AllowedOrigins([]string{"*"})

	// TODO: Need to sort out the port here .. may need to get from environment. Issue when running locally as both FE/BE are on :8080
	if err := http.ListenAndServe(":1234", handlers.CORS(headers, methods, origins)(handler.Router)); err != nil {
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


