package models

import "time"

type Stewardgroup struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Leader       int        `json:"leaderid"`
	DateCreated  *time.Time `json:"datecreated"`
	Status       string     `json:"status"`
	ModifiedDate *time.Time `json:"modifieddate"`
	LeaderName   string     `json:"leadername"`
}

type CreateStewardGroup struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	LeaderId int    `json:"leader" binding:"required"`
}
