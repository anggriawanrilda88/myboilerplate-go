package database

import (
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DbConfig for setting connection configuration
type DbConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

// Database set variable struct fo db
type Database struct {
	*gorm.DB
}

// New connection function call on root go
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
