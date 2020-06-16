package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	IniFile *ini.File
)

func init()  {

	var err error

	IniFile, err = ini.Load("app.ini")

	if err != nil {
		log.Fatalln("Fail to open app.ini:", err)
	}

}
