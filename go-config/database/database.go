package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/devkhatri523/ecom-app/go-config/utils"
)

type Options struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	DatabaseName    string
	Protocol        string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	PARAM           string
}

type Database interface {
	Open(options Options)
	Get() interface{}
	Close()
	Ping() error
}

// build dns means url build of database user:password@protocol(host:port)/dbname
func BuildDns(options Options) (string, error) {
	handleError := func(msg string) (string, error) { return "", errors.New(msg) }
	if utils.IsBlank(options.UserName) {
		return handleError("user name cannot be empty")
	}
	if utils.IsBlank(options.Password) {
		return handleError("password cannot be empty")
	}
	if utils.IsBlank(options.Host) {
		return handleError("host name cannot be empty")
	}
	if options.Port <= 0 {
		return handleError("port cannot be 0 or negative")
	}
	if utils.IsBlank(options.DatabaseName) {
		return handleError("database cannot be empty")
	}
	var protocol string
	fmt.Println(protocol)
	if utils.IsBlank(options.Protocol) {
		protocol = "tcp"
	} else {
		protocol = options.Protocol
	}
	var param string
	if utils.IsBlank(options.PARAM) {
		param = "parseTime=true"
	} else {
		param = options.PARAM
	}
	fmt.Println(param)
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d", options.UserName,
		options.Password, options.Host, options.Port), nil

}
