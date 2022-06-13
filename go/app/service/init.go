package service

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"

	containerName := os.Getenv("MYSQL_CONTAINER_NAME")
	// DsName := "root:root@tcp(vol_mysql:3306)/app?charset=utf8"
	if containerName == "" {
		log.Fatal("MYSQL_CONTAINER_NAME is empty")
	}
	DsName := fmt.Sprintf("root:root@tcp(%s:3306)/app?charset=utf8", containerName)
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	// DbEngine.Sync2(new(model.User))
	fmt.Println("init data base ok")
}
