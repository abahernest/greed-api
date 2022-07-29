package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"errors"

	"github.com/abahernest/greed-api/middleware"
	"github.com/abahernest/greed-api/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ichtrojan/thoth"
)

var (
	Logger, _ = thoth.Init("log")
)

func init(){
	fmt.Println("hello init")
	if err:=godotenv.Load(); err != nil {
		Logger.Log(errors.New(".env File Not Found"))
		log.Fatal(".env file not found")
	}

	current_environment := os.Getenv("CURRENT_ENVIRONMENT")
	connectionString := os.Getenv(fmt.Sprintf("%s_DATABASE_URL", current_environment))
	dbName := os.Getenv(fmt.Sprintf("%s_DB_NAME", current_environment))
	fmt.Printf("loading %s env", current_environment)

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(connectionString))
	if err != nil {
		Logger.Log(errors.New("cannot set mgm config"))
		log.Fatal(err)
	}

	log.Println("Loaded Database")
}
func main() {
	// init()
	fmt.Println("Hello world from Greed-api")
	r := mux.NewRouter()
	router := r.PathPrefix("/api/v1").Subrouter()
	routes.SetupRoutes(router)

	router.Use(middleware.CorsMiddleware)

	port := os.Getenv("PORT")
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),router))
}