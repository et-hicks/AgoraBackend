package microservices


type AgoraService interface {
	Init() error
	Run() error
}