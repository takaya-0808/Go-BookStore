package connecter

import "github.com/jinzhu/gorm"

func SetConnection() *gorm.DB {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	Connect := USER + ":" + PASS + "@" + PROTCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	DB, err := gorm.Open(DBMS, Connect)
	if err != nil {
		panic(err)
	}
	return DB
}
