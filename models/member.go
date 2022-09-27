package models

import "time"

type Member struct {
	Id              int        `json:"id" binding:"required" gorm:"primary_key"`
	Surname         string     `json:"surname"`
	Firstname       string     `json:"firstname"`
	Othernames      string     `json:"othernames"`
	Dob             *time.Time `json:"dateofbirth"`
	Gender          string     `json:"gender"`
	Maritalstatus   string     `json:"martialstatus"`
	Employed        string     `json:"employed"`
	Occupation      string     `json:"occupation"`
	Company         string     `json:"company"`
	Companylocation string     `json:"companylocation"`
	Residence       string     `json:"residence"`
	Mobile          string     `json:"mobile"`
	Email           string     `json:"email"`
	Passport        string     `json:"passport"`
	Datecreated     *time.Time `json:"createddate"`
	Status          string     `json:"status"`
	ModifiedDate    *time.Time `json:"modifieddate"`
	Presbytery      string     `json:"presbytery"`
}

type CreateMember struct {
	Id              int    `json:"id"`
	Surname         string `json:"surname" binding:"required"`
	Firstname       string `json:"firstname" binding:"required"`
	Othernames      string `json:"othernames"`
	Dob             string `json:"dob" binding:"required"`
	Gender          string `json:"gender" binding:"required"`
	MaritalStatus   string `json:"maritalstatus" binding:"required"`
	Employed        string `json:"employed" binding:"required"`
	Occupation      string `json:"occupation" binding:"required"`
	Company         string `json:"company" binding:"required"`
	CompanyLocation string `json:"companylocation" binding:"required"`
	Residence       string `json:"residence" binding:"required"`
	Mobile          string `json:"mobile" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Presbytery      string `json:"presbytery" binding:"required"`
}
