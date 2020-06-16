package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-xb-song/applications/app/pkg/conf"
	"log"
)

var Db *gorm.DB

type Model struct {

}

type dbConf struct {
	dbType       string
	dbLogMode    bool
	host         string
	dbName       string
	user         string
	password     string
	tablePrefix  string
}

// init database connection
func init()  {

	var err error

	cfg := getConfig()

	Db, err = gorm.Open(cfg.dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.user,
		cfg.password,
		cfg.host,
		cfg.dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName;
	}


	Db.LogMode(cfg.dbLogMode)

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

}

func CloseDb()  {
	defer Db.Close()
}

// get database configuration from section 'database' in file 'app.ini'
func getConfig() (dbConf) {

	var cfg dbConf

	sec, err := conf.IniFile.GetSection("database")

	if err != nil {
		log.Fatalln("Fail to get section 'database': ", err)
	}

	cfg.dbType = sec.Key("TYPE").String()
	cfg.dbLogMode, _ = sec.Key("LOG_MODE").Bool()
	cfg.dbName = sec.Key("NAME").String()
	cfg.user = sec.Key("USER").String()
	cfg.password = sec.Key("PASSWORD").String()
	cfg.host = sec.Key("HOST").String()
	cfg.tablePrefix = sec.Key("TABLE_PREFIX").String()

	return cfg

}