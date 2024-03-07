package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RenatoValentim/virtual-bookstore/internal/constants/environments"
	"github.com/RenatoValentim/virtual-bookstore/internal/entities"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type authorDataPostgres struct {
	db *sql.DB
}

func NewAuthorDataPostgres() (*authorDataPostgres, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		viper.GetString(environments.DBHost),
		viper.GetString(environments.DBUser),
		viper.GetString(environments.DBPassword),
		viper.GetString(environments.DBName),
		viper.GetString(environments.DBPort),
		viper.GetString(environments.DBSSLMode),
		viper.GetString(environments.DBTimeZone),
	)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Printf("Failed to connect on database: %v\n", err)
		return nil, err
	}
	return &authorDataPostgres{
		db: db,
	}, nil
}

func (adp *authorDataPostgres) Register(author *entities.Author) error {
	author.ID = uuid.New()
	author.CreatedAt = time.Now().Format(time.RFC3339)

	_, err := adp.db.Exec(`
		INSERT INTO authors (id, name, email, description, created_at) VALUES ($1, $2, $3, $4, $5)
	`, author.ID, author.Name, author.Email, author.Description, author.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
