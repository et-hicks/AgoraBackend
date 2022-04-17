package microservices

import "gorm.io/gorm"

type AgoraService interface {
	Init() error
	Run()
}

type DatabaseController interface {
	RunningInstance() *gorm.DB
}