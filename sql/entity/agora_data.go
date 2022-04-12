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
