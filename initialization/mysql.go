package initialization

import "gorm.io/gorm"

func New() *gorm.DB {

	RDb := gorm.DB{
		Config:       nil,
		Error:        nil,
		RowsAffected: 0,
		Statement:    nil,
	}

	return &RDb
}
