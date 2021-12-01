package models

import (
	"time"
)

type Service struct {
	ID                 uint `gorm:"primaryKey; autoIncrement" json:"id"`
	Name               string
	Type               string
	Slug               string `gorm:"uniqueIndex"`
	Thumbnail          string
	AccessKey          string
	Token              string
	Timestamp          string
	ShortDescription   string
	LongDescription    string
	SpecialInstruction string
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
