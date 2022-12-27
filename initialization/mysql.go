package initialization

import "gorm.io/gorm"

type dbc struct {
	RDbRepo *gorm.DB
	WDbRepo *gorm.DB
}

func New() *gorm.DB {

	RDb := gorm.DB{
		Config:       nil,
		Error:        nil,
		RowsAffected: 0,
		Statement:    nil,
	}

	return &RDb
}
