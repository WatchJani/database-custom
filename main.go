package main

import (
	"fmt"
	"log"
	"root/engine"
)

func main() {
	system := engine.NewDataBase("yovan")

	user, err := system.NewTable("user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user)

	table, err := system.SelectTable("user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)
}
