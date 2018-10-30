package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SqlHandler struct {
	DB *gorm.DB
}

func NewSqlHandler(driver, url string) (*SqlHandler, error) {
	db, err := gorm.Open(driver, url)
	if err != nil {
		return nil, err
	}
	return &SqlHandler{
		DB: db,
	}, nil
}

func (s *SqlHandler) Close() {
	if err := s.DB.Close(); err != nil {
		panic(err)
	}
}

func (s *SqlHandler) Insert(value interface{}) error {
	if err := s.DB.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (s *SqlHandler) Update(value interface{}) error {
	if err := s.DB.Save(value).Error; err != nil {
		return err
	}
	return nil
}

func (s *SqlHandler) SelectById(value interface{}) error {
	if !s.DB.First(value).Limit(1).RecordNotFound() {
		return errors.New("record not found.")
	}
	return nil
}

func (s *SqlHandler) Select(value interface{}) error {
	if err := s.DB.Find(value).Error; err != nil {
		return err
	}
	return nil
}
