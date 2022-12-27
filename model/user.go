package model

import (
	"log"
	"time"
)

type User struct {
	Model
	DepartmentId  int       `db:"department_id" json:"department_id"`
	UserName      string    `db:"username" json:"username"`
	No            string    `db:"no" json:"no"`
	NickName      string    `db:"nickname" json:"nickname"`
	Avatar        *string   `db:"avatar" json:"avatar"`
	Password      string    `db:"password" json:"password"`
	LastLoginTime time.Time `db:"last_login_time" json:"last_login_time"`
	LastLoginIp   string    `db:"last_login_ip" json:"last_login_ip"`
}

func GetUsers() (user User) {
	sql := "SELECT * FROM admin_users WHERE id=?"
	err := db.Get(&user, sql, 42)
	if err != nil {
		log.Fatalln("error:", err)
	}
	return
}
