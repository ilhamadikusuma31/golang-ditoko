package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

func (server *Server) Initialize(nameApp string) {
	fmt.Println("welkom to " + nameApp)
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	err := godotenv.Load()
	errorHandling(err)
	var appConfig = AppConfig{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),
	}

	var server = Server{}
	server.Initialize(appConfig.Name)
	server.Run(appConfig.Port)
}
