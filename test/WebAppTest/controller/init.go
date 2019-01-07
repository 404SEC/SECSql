package controller

import (
	m "github.com/404SEC/SECSql/test/WebAppTest/models"

	go_logger "github.com/phachon/go-logger"
)

var Conf m.Config
var Logger *go_logger.Logger

func init() {
	Conf = m.Conf
	Logger = m.Logger
}
