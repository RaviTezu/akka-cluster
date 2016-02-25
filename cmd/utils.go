package cmd

import (
	"log"

	"github.com/go-ini/ini"
)

//ParseINI - To parse the INI file and return the config.
func ParseINI(filename string) *ini.File {
	conf, err := ini.Load(filename)
	if err != nil {
		log.Panic(err)
	}
	return conf
}

func 
