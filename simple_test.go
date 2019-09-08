package main

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"github.com/lfz757077613/goLearn/utils/myMysql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGoRedis(t *testing.T) {
	Convey("test go-redis set/get", t, func() {
		client := redis.NewClient(&redis.Options{
			Addr: myConf.GetString("redis", "addr", "localhost:6379"),
		})
		err := client.Set("key", "value", 0).Err()
		So(err, ShouldBeNil)
		val, err := client.Get("key").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "value")
	})
}

func TestSqlx(t *testing.T) {
	Convey("test sqlx", t, func() {
		database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lfz")
		So(err, ShouldBeNil)
		err = database.Ping()
		So(err, ShouldBeNil)
		var fruits []myMysql.Fruit
		err = database.Select(&fruits, "select * from fruit limit 1")
		So(err, ShouldBeNil)
		So(len(fruits), ShouldBeGreaterThanOrEqualTo, 0)
	})
}
