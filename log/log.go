package log

import (
	"os"

	"github.com/op/go-logging"
	"github.com/zemirco/papertrail"
)

const format = "%{color}%{time:15:04:05.000} %{level:.4s}%{color:reset} %{message}"

var (
	file = os.Getenv("LOG_FILE")
	log  *logging.Logger
)

func init() {
	if file == "" {
		file = os.DevNull
	} else {
		fileWriter, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		remoteWriter := &papertrail.Writer{
			Network: papertrail.UDP,
			Port:    28644,
			Server:  "logs4",
		}
		logging.SetFormatter(logging.MustStringFormatter(format))
		fileBackend := logging.NewLogBackend(fileWriter, "", 0)
		remoteBackend := logging.NewLogBackend(remoteWriter, "", 0)
		stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
		logging.SetBackend(fileBackend, remoteBackend, stderrBackend)
		log = logging.MustGetLogger("pkmnrequiem-cluster")
	}
}

func Critical(msg string, args ...interface{}) {
	log.Criticalf(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

func Fatal(err error) {
	log.Error(err.Error())
}

func Info(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func Notice(msg string, args ...interface{}) {
	log.Noticef(msg, args...)
}

func Warning(msg string, args ...interface{}) {
	log.Warningf(msg, args...)
}
