package libs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Addr string
	User string
}

func ConnectMySQL() *sql.DB {
	DriverName := "mysql"
	Dsn := "root:secret@tcp(172.20.1.5:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open(DriverName, Dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
