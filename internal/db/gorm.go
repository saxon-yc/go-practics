package dbsvc

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewGormContext initializes the MySQL database connections.
func NewGormContext(port int, host, user, password, dbname string) *gorm.DB {
	gormConfig := gorm.Config{}
	gormConfig.NamingStrategy = schema.NamingStrategy{
		SingularTable: true,
	}

	// "username:password@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gormConfig)

	if err != nil {
		s := fmt.Sprintf("can't connect to database:%s,error:%s\n", dsn, err)
		log.Fatal(s)
	}
	sqlDb, err := db.DB()
	if err != nil {
		s := fmt.Sprintf("get connection pool failed, error: %s\n", err)
		log.Fatal(s)
	}
	sqlDb.SetConnMaxLifetime(8 * time.Second)
	sqlDb.SetMaxOpenConns(20)

	return db
}
