package main

import (
	"Taskie/cfg"
	"fmt"
)

func main() {
	cfg, err := cfg.Load()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("config: %+v\n", *cfg)
}
