package models

import "time"

type User struct {
	Id             int        `json:"id" binding:"required" gorm:"primary_key"`
	Username       string     `json:"username"`
	Password       string     `json:"-"`
	CreatedBy      int        `json:"created_by"`
	Datecreated    *time.Time `json:"created_date"`
	Msisdn         string     `json:"mobile"`
	Status         string     `json:"status"`
	Lastlogindate  *time.Time `json:"last_login"`
	Lastupdatetime *time.Time `json:"last_updated"`
	Firstname      string     `json:"firstname"`
	Lastname       string     `json:"lastname"`
}

type CreateUser struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Msisdn    string `json:"mobile" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePassword struct {
	Id          int    `json:"id" binding:"required"`
	Oldpassword string `json:"oldpassword" binding:"required"`
	Newpassword string `json:"newpassword" binding:"required"`
}
