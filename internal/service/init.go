package service

import (
	"github.com/apsdehal/go-logger"
	"massimple.com/wallet-controller/internal/service/gpIssuer"
	"os"
)

var log *logger.Logger

func Init() {
	var err error
	log, err = logger.New("Service", 1, os.Stdout)
	if err != nil{
		panic(err)
	}
	SMSInit()
	gpIssuer.GPInit()
}
