package main

import (
	"fmt"
	"log"
	"root/engine"
)

func main() {
	system := engine.NewDataBase("yovan")

	user, err := system.CreateTable("user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user)

	user, err = system.SelectTable("user") //i want to send error to our new method
	if err != nil {
		log.Println(err)
	}

	// user.AddIndex(index.TRIE)

	fmt.Println(user)
}
