package schemas

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Opening struct {
	gorm.Model
	Role string 
	Company string
	Location string
	Remote bool
	Link string
	Salary int64
}

