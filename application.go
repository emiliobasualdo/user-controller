package main

import (
	_ "github.com/mattn/go-sqlite3"
	"massimple.com/wallet-controller/internal/persistence"
	"massimple.com/wallet-controller/internal/service"
	"massimple.com/wallet-controller/internal/utils"
	"massimple.com/wallet-controller/internal/webapp"
	"math/rand"
	"os"
	"time"
)


func main() {
	value, exists := os.LookupEnv("ENV")
	if !exists {
		panic("Environment variable ENV must be set to PRO or DEV or TEST")
	}
	rand.Seed(time.Now().UnixNano())
	utils.EnvInit(value)
	persistence.Init()
	service.Init()
	webapp.Serve()
}




