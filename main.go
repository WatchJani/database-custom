package main

import (
	"fmt"
	"log"
	"root/engine"
	"root/index"
)

type DataFlow struct {
	*engine.DataBase
}

func main() {
	system := engine.NewDataBase("yovan")

	user, err := system.CreateTable("user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user)

	//lov level func
	user, err = system.SelectTable("user") //i want to send error to our new method
	if err != nil {
		log.Println(err)
	}

	user.AddIndex(index.TRIE)

	fmt.Println(user)

	//height level func
	system.AddIndex("user", index.TRIE)
}
