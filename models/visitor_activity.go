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

func (m *Model) GetVisitorActivities(VisitorActivity *[]VisitorActivity) (err error) {
	err = m.DBConn.Find(VisitorActivity).Error
	if err != nil {
		return err
	}
	return nil
}