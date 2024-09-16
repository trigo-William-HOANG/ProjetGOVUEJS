package tools

import (
	log "github.com/sirupsen/logrus"
)

type AppDetails struct {
	Id          int
	Title       string
	Description string
	Link        string
	Image       string
}

type DatabaseInterface interface {
	GetAppDetails(Id int) *AppDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &fausseDB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
