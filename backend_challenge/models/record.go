package models

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model `json:"-"`
	ID uint `json:"id"`
	Age           int
	Workclass     Workclass
	Fnlwgt        int
	Education     string
	EducationNum  int
	MaritalStatus string
	Occupation    string
	Relationship  string
	Race          string
	Sex           string
	CapitalGain   int
	CapitalLoss   int
	HoursPerWeek  int
	NativeCountry string
	Income        string
}

type APIRecord struct {
	ID uint 
	Age           int
	Education     string
	MaritalStatus string
	Occupation    string
	Income        string
	HoursPerWeek int
}
