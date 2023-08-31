package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rtanx/gostarter/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connInstance *gorm.DB
var once sync.Once

func newConn(connCfg *postgres.Config) *gorm.DB {
	db, err := gorm.Open(postgres.New(*connCfg), &gorm.Config{
		Logger: NewDBLogger(),
	})
	if err != nil {
		log.Fatalf("cannot initialized database connection: %v", err)
		return nil
	}
	return db
}

func GetConn() *gorm.DB {
	once.Do(func() {
		log.Println("initialized database connection")
		connInstance = newConn(connConfig())
	})
	return connInstance
}

func connConfig() *postgres.Config {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.DBUser(), env.DBPassword(),
		env.DBHost(), env.DBPort(),
		env.DBName(), env.DBSslMode(),
	)
	return &postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}
}

func NewDBLogger() logger.Interface {
	dbgLevel := env.DBDebug()
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(dbgLevel),
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
}
