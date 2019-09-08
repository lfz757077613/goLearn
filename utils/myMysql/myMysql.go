package myMysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"time"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", myConf.GetString("mysql", "dataSourceName", "root:root@tcp(127.0.0.1:3306)/lfz"))
	if err != nil {
		panic(err)
	}
	database.SetMaxIdleConns(myConf.GetInt("mysql", "maxIdleConns", 10))
	database.SetMaxOpenConns(myConf.GetInt("mysql", "maxOpenConns", 10))
	database.SetConnMaxLifetime(time.Second * time.Duration(myConf.GetInt64("mysql", "maxLifetime", 10)))
	err = database.Ping()
	if err != nil {
		panic(err)
	}
	Db = database
}
