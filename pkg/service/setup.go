package service

import "github.com/apsdehal/go-logger"

var log *logger.Logger

func Init(_log *logger.Logger) {
	log = _log
}
