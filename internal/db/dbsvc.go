package dbsvc

import (
	"context"
	"fmt"
	"go-practics/config"
	"go-practics/internal/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
	redisUrl := fmt.Sprintf("%s:%d", config.Host, config.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	ctx := context.Background()
	s, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Init redis filed: %v. \n", err)
	}
	log.Printf("Init redis success: %v. \n", s)

	return rdb, ctx
}
func NewDbServer() {
	newMysql()
	newRedis()

}
