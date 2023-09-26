package main

import (
	emiHandler "EMI/controller/handler"
	emiRepo "EMI/controller/repository"
	emiService "EMI/controller/service"
	"database/sql"
	"io"
	"log"
	"os"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(err)
	}
}

func init() {
	gotenv.Load()
}

func setupLogOutput() {
	f, _ := os.Create("gin.Log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	port := viper.GetString("portlocal.port")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer db.Close()

	router := gin.New()

	Repo := emiRepo.CreateEmiRepoMysqlImpl(db)
	Service := emiService.CreateEmiServiceImpl(Repo)
	emiHandler.CreateEmiHandler(router, Service)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
