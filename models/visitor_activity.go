package models

import (
	"time"
)

type VisitorActivity struct {
	ID           uint    `gorm:"primaryKey; autoIncrement"`
	SessionID    string  `json:"session_id"`
	Visitor      Visitor `gorm:"foreignKey:SessionID"` // `VisitorActivity` belongs to `Visitor`, `SessionID` is the foreign key
	ServiceID    uint    `json:"service_id"`
	Service      Service `gorm:"foreignKey:ServiceID"` // `VisitorActivity` belongs to `Service`, `ServiceID` is the foreign key
	Completeness int     `json:"completeness"`
	CreatedAt    time.Time
}

type VisitorActivities struct {
	Email			string
	FullName		string
	Company			string
	JobTitle		string
	Industry		string
	Name			string
	Type			string
	Completeness 	int
	CreatedAt		time.Time
}

func (m *Model) GetVisitorActivities(VisitorActivities *[]VisitorActivities) (err error) {
	err = m.DBConn.Table("visitors AS v").
	Select("va.created_at, v.email, v.full_name, v.company, v.job_title, v.industry, s.name, s.type, va.completeness").
	Joins("JOIN visitor_activities va ON va.session_id = v.session_id").
	Joins("JOIN services s ON s.id = va.service_id").
	Find(VisitorActivities).Error

	if err != nil {
		return err
	}
	
	return nil
}
