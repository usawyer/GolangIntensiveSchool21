package db

import (
	"day06/internal/service"
	"day06/models"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"os"
	"time"
)

type pgClient struct {
	db *gorm.DB
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func New(zapLogger *zap.Logger) service.Store {
	connectionParams := map[string]string{
		"host":     getEnv("DB_HOST", "localhost"),
		"user":     getEnv("POSTGRES_USER", "postgres"),
		"password": getEnv("POSTGRES_PASSWORD", "postgres"),
		"dbname":   getEnv("POSTGRES_DB", "day06"),
		"port":     getEnv("DB_PORT", "5432"),
		"sslmode":  "disable",
		"TimeZone": "Asia/Novosibirsk",
	}
	gormLogger := zapgorm2.New(zapLogger)
	var dsn string

	for key, value := range connectionParams {
		dsn += fmt.Sprintf("%s=%s ", key, value)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			zapLogger.Warn("Error open", zap.Error(err))
			continue
		}
		db = db.Debug()
		err = db.AutoMigrate(&models.Articles{})
		if err != nil {
			zapLogger.Error(err.Error())
		} else {
			zapLogger.Info("migrate ok")
			//var count int64
			//db.Model(&models.Author{}).Count(&count)
			//if count == 0 {
			//	err := fillTableFromSQLFile(db, "res/sql-migrations/users.sql")
			//	err = fillTableFromSQLFile(db, "res/sql-migrations/posts.sql")
			//
			//	if err != nil {
			//		log.Fatal(err)
			//	}
			//}
		}
		return &pgClient{db: db}
	}
	zapLogger.Fatal("Error open db")
	return nil
}

func (p *pgClient) AddArticle(article models.Articles) error {
	res := p.db.Create(&article)
	return res.Error
}

func (p *pgClient) GetArticle(id int) (models.Articles, error) {
	var article models.Articles
	res := p.db.First(&article, id)
	return article, res.Error
}

func (p *pgClient) GetArticles(params models.SearchParams) ([]models.Articles, error) {
	var articles []models.Articles
	res := p.db.Offset(params.Offset).Limit(params.Limit).Find(&articles)
	return articles, res.Error
}

func (p *pgClient) RemoveArticle(id int) error {
	res := p.db.Delete(&models.Articles{}, id)
	if res.Error == nil && res.RowsAffected != 1 {
		return errors.New("article with such id doesn't exist")

	}
	return res.Error
}
