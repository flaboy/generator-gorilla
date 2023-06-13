package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/revel/revel"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

// DB Gorm.
var (
	DB      *gorm.DB
	gormLog = revel.AppLog
)

func init() {
	revel.RegisterModuleInit(func(m *revel.Module) {
		gormLog = m.Log
	})
}

type DbInfo struct {
	DbDriver   string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
}

func InitDBWithParameters(params DbInfo) {
	dbInfo := ""
	var driver gorm.Dialector
	switch params.DbDriver {
	default:
		dbInfo = fmt.Sprintf(params.DbHost)
		driver = sqlite.Open(dbInfo)
	case "postgres":
		dbInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", params.DbHost, params.DbPort, params.DbUser, params.DbName, params.DbPassword)
		driver = postgres.Open(dbInfo)
	case "mysql":
		dbInfo = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", params.DbUser, params.DbPassword, params.DbHost, params.DbPort, params.DbName)
		driver = mysql.Open(dbInfo)
	}

	config := &gorm.Config{}
	config.NamingStrategy = schema.NamingStrategy{
		SingularTable: revel.Config.BoolDefault("db.singulartable", false),
	}

	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		gormLog.Fatal("sql.Open failed", "error", err)
	}
	DB = db
}

func InitDB() {
	params := DbInfo{}
	params.DbDriver = revel.Config.StringDefault("db.driver", "sqlite3")
	params.DbHost = revel.Config.StringDefault("db.host", "localhost")

	switch params.DbDriver {
	case "postgres":
		params.DbPort = revel.Config.IntDefault("db.port", 5432)
	case "mysql":
		params.DbPort = revel.Config.IntDefault("db.port", 3306)
	case "sqlite3":
		if params.DbHost == "localhost" {
			params.DbHost = "/tmp/app.db"
		}
	}

	params.DbUser = revel.Config.StringDefault("db.user", "default")
	params.DbPassword = revel.Config.StringDefault("db.password", "")
	params.DbName = revel.Config.StringDefault("db.name", "default")

	InitDBWithParameters(params)
}
