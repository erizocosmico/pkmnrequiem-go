package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/services"
	"gopkg.in/mgo.v2"
)

var (
	bindAddress = os.Getenv("RUN_ADDRESS")
	dbName      = os.Getenv("MONGO_DB")
	dbURI       = os.Getenv("MONGO_URI")
)

func main() {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	v1 := router.Group("v1")
	services.Services{
		services.NewAccountService(db),
		services.NewBattleService(db),
	}.Register(v1)
	log.Fatal(router.Run(bindAddress))
}

func dbConnection() (*mgo.Database, error) {
	session, err := mgo.Dial(dbURI)
	if err != nil {
		return nil, err
	}
	return session.DB(dbName), nil
}
