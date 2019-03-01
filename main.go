package main

import (
	"fmt"

	db "github.com/salihkemaloglu/goTest/db"
	repos "github.com/salihkemaloglu/goTest/repos"
)

func main() {

	repo := db.Item{Name: "hey"}
	var e repos.BaseRepository = repo
	var items, _ = e.FindByID()

	fmt.Println(items)

}
