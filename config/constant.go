package config

const (
	Host = "139.198.41.145"

	MysqlUsername = "root"
	MysqlPasswd   = "QweRt123$"
	MysqlPort     = 3306
	AutoMigrate   = true
	MysqlDbName   = "m_db"

	PsqlUsername = "root"
	PsqlPasswd   = "QweRt123$"
	PsqlPort     = 5432
	PsqlDbName   = "m_db"

	RedisPort = 6379

	KafkaTopic = "my-topic"
	KafkaPort  = 9092
)
