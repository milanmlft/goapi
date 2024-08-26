package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections; what can the database return?
type LoginDetails struct {
	AuthToken string
	Username  string
}

type flavourDetails struct {
	Name     string
	Quantity int64
}

// Slice of flavours
type Flavours []flavourDetails

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetFlavours() *Flavours
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
