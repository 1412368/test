package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Config DB info
type Config struct {
	user, pass, url, db, maxIdle, maxConn string
}

//InitDatabase is use to create db connection
func InitDatabase() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	config := &Config{
		user:    beego.AppConfig.String("MYSQL_USER"),
		pass:    beego.AppConfig.String("MYSQL_PASSWORD"),
		url:     beego.AppConfig.String("MYSQL_URL"),
		db:      beego.AppConfig.String("MYSQL_DB"),
		maxIdle: beego.AppConfig.String("MAX_IDLE"),
		maxConn: beego.AppConfig.String("MAX_CONN"),
	}
	var connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8", config.user, config.pass, config.db)
	var maxIdle, _ = strconv.Atoi(config.maxIdle)
	var maxConn, _ = strconv.Atoi(config.maxConn)
	log := logs.GetLogger()
	log.Println(connectionString)
	err := orm.RegisterDataBase("default", "mysql", connectionString, maxIdle, maxConn)
	if err != nil {
		panic(err)
	}
	orm.DefaultTimeLoc = time.UTC

	defaultDB, err := orm.GetDB("default")
	if err != nil {
		panic(err)
	}

	defaultDB.SetConnMaxLifetime(300 * time.Second)
	orm.Debug = true
	err = orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
}
