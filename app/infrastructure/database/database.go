package database

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Redis *redis.Client

// DbConfig for setting connection configuration
type DbConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

// DbConfig for setting connection configuration
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// Database set variable struct fo db
type Database struct {
	*gorm.DB
}

// New connection function call on root go
func NewRedis(config *RedisConfig) (err error) {
	// add redis connection
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err = Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("failed to connect redis:", err.Error())
		return
	}

	return
}

func New(config *DbConfig) (*Database, error) {
	var err error
	switch strings.ToLower(config.Driver) {
	case "mysql":
		dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgresql", "postgres":
		dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) + "?database=" + config.Database
		DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	}
	return &Database{DB}, err
}
