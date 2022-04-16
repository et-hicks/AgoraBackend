package main

import (
	"fmt"
	"github.com/admin-agora/backend/src/microservices/posting"
)

func main() {
	var post posting.PostProcessing

	initErr := post.Init()
	if initErr != nil {
		return
	}

	runErr := post.Run()
	if runErr != nil {
		return
	}

	fmt.Println("shutting down")

}
