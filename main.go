package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/log"
	"github.com/jrevillas/pkmnrequiem-go/services"
	"gopkg.in/mgo.v2"
)

var (
	bindAddress = os.Getenv("RUN_ADDRESS")
	dbName      = os.Getenv("MONGO_DB")
	dbURI       = os.Getenv("MONGO_URI")
)

func main() {
	log.Notice("pkmnRequiem 0.1")
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
	log.Debug("Establishing database connection...")
	start := time.Now()
	session, err := mgo.Dial(dbURI)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	log.Debug("Connected to the database (%.2fms)", time.Since(start).Seconds()*1000)
	return session.DB(dbName), nil
}
