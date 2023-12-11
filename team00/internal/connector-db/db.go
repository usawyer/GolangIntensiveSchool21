package connector_db

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func New(logger *zap.Logger) *gorm.DB {
	connectionParams := map[string]string{
		"host":     getEnv("DB_HOST", "localhost"),
		"user":     getEnv("POSTGRES_USER", "postgres"),
		"password": getEnv("POSTGRES_PASSWORD", "postgres"),
		"dbname":   getEnv("POSTGRES_DB", "team00"),
		"port":     getEnv("DB_PORT", "5432"),
		"sslmode":  "disable",
		"TimeZone": "Asia/Novosibirsk",
	}
	var dsn string
	for key, value := range connectionParams {
		dsn += fmt.Sprintf("%s=%s ", key, value)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Warn("Error open", zap.Error(err))
			continue
		}
		return db
	}
	logger.Fatal("Error open db")
	return nil
}
