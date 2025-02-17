package main

import (
	"Taskie/cfg"
	"Taskie/db"
	"fmt"
)

func main() {
	cfg, err := cfg.Load()
	if err != nil {
		fmt.Println(err)
	}

	db, err := db.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}
