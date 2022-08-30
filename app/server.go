package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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

type DbConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func (server *Server) Initialize(nameApp string, dbConfig DbConfig) {
	fmt.Println("welkom to " + nameApp)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/jakarta", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	var err error
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errorHandling(err)

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

	var dbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	var server = Server{}
	server.Initialize(appConfig.Name, dbConfig)
	server.Run(appConfig.Port)
}
