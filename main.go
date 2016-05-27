package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/services"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
)

var (
	bindAddress  = os.Getenv("RUN_ADDRESS")
	dbName       = os.Getenv("MONGO_DB")
	dbURI        = os.Getenv("MONGO_URI")
	format       = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{level:.4s} %{id:03x}%{color:reset} %{message}`)
	log          = logging.MustGetLogger("logger")
	logBackend   = logging.NewLogBackend(os.Stdin, "", 0)
	logFormatter = logging.NewBackendFormatter(logBackend, format)
)

func main() {
	logging.SetBackend(logFormatter)
	log.Notice("pkmnRequiem 0.1")
	log.Notice("Conectando con la base de datos...")
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
	start := time.Now()
	session, err := mgo.Dial(dbURI)
	if err != nil {
		return nil, err
	}
	log.Noticef("Base de datos conectada (%.2fms)", time.Since(start).Seconds()*1000)
	return session.DB(dbName), nil
}
