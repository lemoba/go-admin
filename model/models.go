package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lemoba/go-admin/pkg/setting"
	"log"
	"time"
)

var db *sqlx.DB

type Model struct {
	Id        int       `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func init() {
	var (
		dbConnection string
		dbHost       string
		dbPort       int
		dbName       string
		dbUserName   string
		dbPassword   string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	dbConnection = sec.Key("DB_CONNECTION").String()
	dbHost = sec.Key("DB_HOST").String()
	dbPort, _ = sec.Key("DB_PORT").Int()
	dbName = sec.Key("DB_DATABASE").String()
	dbUserName = sec.Key("DB_USERNAME").String()
	dbPassword = sec.Key("DB_PASSWORD").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", dbUserName, dbPassword, dbHost, dbPort, dbName)
	db, err = sqlx.Open(dbConnection, dsn)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("连接数据库成功")
}
