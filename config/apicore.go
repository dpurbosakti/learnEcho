package config

/*****************************************************************************\
The idea is to place all the needs for a web application to be running in one
place. All the user to do is set the configuration and the handlers since it
is main properties of a website.
\*****************************************************************************/

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type app struct {
	Code            string
	Name            string
	Env             string
	Version         string
	HttpConf        *httpConf
	DbConf          *dbConf
	MailerConf      *mailerConf
	RateLimiterConf *rateLimiterConf
	LoggerConf      *loggerConf
}

// Exported package vars for easier resolving
var Self *app // the core itself
var DB *gorm.DB
var Logger *zap.Logger

// app starter
func Run(routerIn *echo.Echo, code string) {
	Self = &app{}

	// Intantiate all the conf
	Self.Code = code
	Self.HttpConf = &httpConf{}
	Self.DbConf = &dbConf{}
	Self.MailerConf = &mailerConf{}
	Self.LoggerConf = &loggerConf{}

	// init all manually to make it clear of what's happening
	Self.initConfig()
	Self.initLogger()
	Self.initDb()
	// // a.initMailer()
	Self.initExtCall()
	Self.initHttp(routerIn)
}
