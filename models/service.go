package models

import (
	"time"
)

type Service struct {
	ID                 uint `gorm:"primaryKey; autoIncrement" json:"id" form:"id"`
	Name               string `form:"name"`
	Type               string `form:"type"`
	Slug               string `gorm:"uniqueIndex" form:"slug"`
	Thumbnail          string `form:"thumbnail"`
	AccessKey          string `form:"access_key"`
	Token              string `form:"token"`
	Timestamp          string `form:"timestamp"`
	ShortDescription   string `form:"short_description"`
	LongDescription    string `form:"long_description"`
	SpecialInstruction string `form:"special_instruction"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (m *Model) CreateService(Service *Service) (err error) {
	err = m.DBConn.Create(Service).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetServices(Service *[]Service) (err error) {
	err = m.DBConn.Find(Service).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetService(Service *Service, id string) (err error) {
	err = m.DBConn.Where("id = ?", id).First(Service).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) UpdateService(Service *Service) (err error) {
	m.DBConn.Save(Service)
	return nil
}

func (m *Model) DeleteService(Service *Service, id string) (err error) {
	m.DBConn.Where("id = ?", id).Delete(Service)
	return nil
}
