package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/services"
	"gopkg.in/mgo.v2"
)

var (
	mongoDatabase = os.Getenv("MONGO_DB")
	mongoURI      = os.Getenv("MONGO_URI")
	address       = os.Getenv("RUN_ADDRESS")
)

func main() {
	db, err := databaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	v1 := router.Group("v1")
	services.Services{
		services.NewAccountService(db),
		services.NewBattleService(db),
	}.Register(v1)
	log.Fatal(router.Run(address))
}

func databaseConnection() (*mgo.Database, error) {
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		return nil, err
	}
	return session.DB(mongoDatabase), nil
}
