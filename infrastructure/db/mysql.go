package db

import (
	"fmt"
	"test-api/infrastructure/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlClient struct {
	*gorm.DB
}

// db is the postgres instance
var db *MysqlClient

// Get returns the default postgresClient currently in use
func Get() *MysqlClient {
	return db
}

// Connect database, must call once before server boot to Get() the db instance
func Connect() error {
	cnfg := config.Get().Database
	userName := cnfg.Username
	pass := cnfg.Password
	host := cnfg.Host
	port := cnfg.Port
	dbname := cnfg.Name

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", userName, pass, host, port, dbname)



	// open connection to postgres db
	instance, err := gorm.Open("mysql", dbSource)
	if err != nil {
		return err
	}

	// connection pool settings
	if cnfg.MaxLifeTime != 0 {
		instance.DB().SetConnMaxLifetime(cnfg.MaxLifeTime * time.Second)
	}
	if cnfg.MaxIdleConn != 0 {
		instance.DB().SetMaxIdleConns(cnfg.MaxIdleConn)
	}
	if cnfg.MaxOpenConn != 0 {
		instance.DB().SetMaxOpenConns(cnfg.MaxOpenConn)
	}

	instance.LogMode(cnfg.Debug)
	db = &MysqlClient{DB: instance}
	return nil
}
