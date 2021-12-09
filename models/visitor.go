package models

import (
	"fmt"
	"time"
)

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

func (m *Model) GetVisitors(Visitor *[]Visitor) (err error) {
	err = m.DBConn.Find(Visitor).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetAMLPEPVisitors() (ListVisitor []Visitor, err error) {
	ListVisitor = []Visitor{}

	var service Service
	err = m.GetServiceBySlug(&service, "aml-pep")
	if err != nil {
		fmt.Println(err)
		return ListVisitor, err
	}

	var visitorActivities []VisitorActivity
	err = m.DBConn.Where("service_id = ?", service.ID).Where("completeness = ?", 100).Find(&visitorActivities).Error
	if err != nil {
		return ListVisitor, err
	}

	for i := 0; i < len(visitorActivities); i++ {
		visitorActivity := visitorActivities[i]
		var visitor Visitor
		m.DBConn.Where("session_id = ?", visitorActivity.SessionID).Find(&visitor)
		ListVisitor = append(ListVisitor, visitor)
	}
	return ListVisitor, nil
}
