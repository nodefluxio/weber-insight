package models

import (
	"time"
)

type Feedback struct {
	ID                uint `gorm:"primary_key"`
	VisitorActivityID uint
	VisitorActivity   VisitorActivity `gorm:"foreignKey:VisitorActivityID"`
	Rating            uint            `json:"rating" validate:"required,min=1,max=5"`
	Comment           string          `json:"comment" validate:"required,min=20,max=255"`
	CreatedAt         time.Time
}

type FeedbackResultView struct {
	Time        time.Time
	ServiceName string
	Email       string
	FullName    string
	Rating      uint
	Comment     string
}

func (m *Model) GetAllFeedback(Feedback *[]Feedback) (err error) {
	err = m.DBConn.Find(Feedback).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetAllFeedbackToView() (ListFeedbackResultView []FeedbackResultView, err error) {
	var allFeedback []Feedback
	m.GetAllFeedback(&allFeedback)

	for i := 0; i < len(allFeedback); i++ {
		feedback := allFeedback[i]
		var visitorActivity VisitorActivity
		m.DBConn.Where("id = ?", feedback.VisitorActivityID).Find(&visitorActivity)

		var service Service
		m.DBConn.Where("id = ?", visitorActivity.ServiceID).Find(&service)

		var visitor Visitor
		m.DBConn.Where("session_id = ?", visitorActivity.SessionID).Find(&visitor)

		var feedbackResultView FeedbackResultView
		feedbackResultView.Time = feedback.CreatedAt
		feedbackResultView.Comment = feedback.Comment
		feedbackResultView.Email = visitor.Email
		feedbackResultView.Rating = feedback.Rating
		feedbackResultView.ServiceName = service.Name
		feedbackResultView.FullName = visitor.FullName
		ListFeedbackResultView = append(ListFeedbackResultView, feedbackResultView)
	}

	return ListFeedbackResultView, nil
}
