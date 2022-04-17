package microservices

import "net/http"

type AgoraService interface {
	Init() error
	ProcessRequest(w http.ResponseWriter, r *http.Request)
}