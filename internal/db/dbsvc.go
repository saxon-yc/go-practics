package dbsvc

import (
	"context"
	"fmt"
	"log"

	// Importing the MySQL driver for its side-effects (e.g., to register the driver)
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"go-practics/config"
	"go-practics/internal/model"
)

func newMysql() *gorm.DB {
	gdb := NewGormContext(config.MysqlPort, config.Host, config.MysqlUsername, config.MysqlPasswd, config.MysqlDbName)

	if config.AutoMigrate {
		err := gdb.AutoMigrate(&model.UserDB{})
		if err != nil {
			log.Fatalf("migrate table error[%s] exited \n", err)
		}
		log.Print("Init database success \n")
	}

	return gdb
}

func newRedis() (*redis.Client, context.Context) {
	redisURL := fmt.Sprintf("%s:%d", config.Host, config.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	ctx := context.Background()
	s, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Init redis failed: %v. \n", err)
	}
	log.Printf("Init redis success: %v. \n", s)

	return rdb, ctx
}

// NewDbServer initializes the MySQL and Redis database connections.
func NewDbServer() {
	newMysql()
	newRedis()
}
