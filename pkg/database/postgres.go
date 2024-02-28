package database

import (
	"fmt"
	"log"
	"notes-api/config"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Postgres(conf *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.DbHost,
		conf.DbUser,
		conf.DbPass,
		conf.DbName,
		conf.DbPort,
	)

	gormConf := &gorm.Config{}

	if !conf.DbDebug {
		gormConf.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent,
				Colorful:      true,
			},
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConf)
	if err != nil {
		return nil, err
	}

	if conf.DbDebug {
		return db.Debug(), nil
	}
	return db, err
}
