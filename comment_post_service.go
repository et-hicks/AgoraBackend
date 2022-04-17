package main

import (
	"github.com/admin-agora/backend/src/microservices/posting"
	"log"
	"net/http"
)

func CommentPostService(w http.ResponseWriter, r *http.Request) {
	var commentPosting posting.CommentPostProcessing

	initErr := commentPosting.Init()
	if initErr != nil {
		return
	}

	commentPosting.ProcessRequest(w, r)

	log.Println("things have finished")

}
