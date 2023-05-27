package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)


type Person struct{
	gorm.Model
	Given_Name string
	Middle_Name string
	Family_Name string
	Email_Array []string
	Email_Index int
	Phone_Array []string
	Phone_Index int
	AddressID int
	Address Address
	DOB int
	OYL int
	OLP int
	LPC int
	YLC int
	FNDP int
	ONDP int
	GPC int
	GPO int
	CPC int
	PCPO int
  }