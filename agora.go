package main

import (
	"fmt"
	"github.com/admin-agora/backend/protobufs/protobufs"
)

func main() {

	// this is not really for anything but to act as a github file making sure I set up the depencencies correcly
	// ok good
	person := protobufs.Person{
		Name:        "",
		Id:          0,
		Email:       "",
		Phones:      nil,
		LastUpdated: nil,
	}

	fmt.Println(person)

}
