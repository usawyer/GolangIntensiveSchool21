package connector_db

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"

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

func New() (*gorm.DB, error) {
	connectionParams := map[string]string{
		"host":     getEnv("DB_HOST", "localhost"),
		"user":     getEnv("POSTGRES_USER", "postgres"),
		"password": getEnv("POSTGRES_PASSWORD", "postgres"),
		"dbname":   getEnv("POSTGRES_DB", "day06"),
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
			log.Println("Error open", err)
			continue
		}
		return db, nil
	}
	return nil, errors.New("Error open db")
}
