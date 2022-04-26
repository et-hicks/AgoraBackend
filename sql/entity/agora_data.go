package entity

import "gorm.io/gorm"

type AgoraData interface {
	LoadForCode() error
	UnloadForDatabase() error
}

type AgoraSQL interface {
	CreateTable(db *gorm.DB) error
	UpdateTable(db *gorm.DB, sqlUpdate string) error
}

// AgoraSQLConstraints TODO: merge this interface with the one above
type AgoraSQLConstraints interface {
	AddConstraints(db *gorm.DB) error
	DeleteConstraints(db *gorm.DB) error
}
