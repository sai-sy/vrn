package models

type Address struct{
	gorm.Model 
	Unit string
	Street_Number string
	Street string
	City string
	District int
	Country []string
	Postal_Code int
  }