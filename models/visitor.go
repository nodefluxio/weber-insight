package models

import "time"

type Visitor struct {
	SessionID string    `gorm:"primary_key" json:"session_id"`
	Email     string    `json:"email" validate:"required,email"`
	FullName  string    `json:"full_name" validate:"required,min=2,max=255"`
	Company   string    `json:"company" validate:"required,min=2,max=255"`
	JobTitle  string    `json:"job_title" validate:"required,min=2,max=255"`
	Industry  string    `json:"industry" validate:"required,min=2,max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
