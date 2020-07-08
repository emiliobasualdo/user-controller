package main

import (
	"github.com/apsdehal/go-logger"
	_ "github.com/mattn/go-sqlite3"
	"massimple.com/wallet-controller/pkg/persistence"
	"massimple.com/wallet-controller/pkg/service"
	"massimple.com/wallet-controller/pkg/webapp"
	"os"
)

var log *logger.Logger

func main() {
	_log, err := logger.New("Wallet", 1, os.Stdout)
	log = _log
	if err != nil {
		panic(err)
	}
	persistence.Init(log)
	service.Init(log)
	webapp.Serve(log)
}


