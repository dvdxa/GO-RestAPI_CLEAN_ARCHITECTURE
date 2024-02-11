package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(dbParams *config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=%s",
		dbParams.Server, dbParams.Port, dbParams.User, dbParams.Password, dbParams.Database, dbParams.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		log.Fatal("failed to init DB ", err)
	}

	return db
}
