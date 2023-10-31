package services

import (
	"errors"
	"job-portal-api/internal/models"

	"gorm.io/gorm"
)

type Conn struct {

	// db is an instance of the SQLite database.
	db *gorm.DB
}

func NewConn(dbInstance *gorm.DB) (*Conn, error) {
	if dbInstance == nil {
		return nil, errors.New("provide the database instance")
	}
	return &Conn{db: dbInstance}, nil
}
func (s *Conn) AutoMigrate() error {

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err := s.db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
