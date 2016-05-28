package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/services"
	"github.com/op/go-logging"
	"github.com/zemirco/papertrail"
	"gopkg.in/mgo.v2"
)

var (
	bindAddress  = os.Getenv("RUN_ADDRESS")
	dbName       = os.Getenv("MONGO_DB")
	dbURI        = os.Getenv("MONGO_URI")
	format       = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{level:.4s}%{color:reset} %{message}`)
	log          = logging.MustGetLogger("logger")
	logBackend   = logging.NewLogBackend(os.Stdin, "", 0)
	logFormatter = logging.NewBackendFormatter(logBackend, format)
)

func main() {
	w := papertrail.Writer{
		Network: papertrail.UDP,
		Port:    28644,
		Server:  "logs4",
	}
	logRemote := logging.NewLogBackend(&w, "", 0)
	logRemoteFormatter := logging.NewBackendFormatter(logRemote, format)
	logging.SetBackend(logFormatter, logRemoteFormatter)
	log.Notice("pkmnRequiem 0.1")
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	v1 := router.Group("v1")
	services.Services{
		services.NewAccountService(db),
		services.NewBattleService(db, log),
	}.Register(v1)
	log.Fatal(router.Run(bindAddress))
}

func dbConnection() (*mgo.Database, error) {
	log.Debug("Conectando con la base de datos...")
	start := time.Now()
	session, err := mgo.Dial(dbURI)
	if err != nil {
		return nil, err
	}
	log.Debugf("Base de datos conectada (%.2fms)", time.Since(start).Seconds()*1000)
	return session.DB(dbName), nil
}
