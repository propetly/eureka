package main

import (
	"api/internal/config"
	"fmt"
)

func main() {

	cfg := config.LoadEnv()

	fmt.Println("Init project")
	//fmt.Println(cfg)
	//fmt.Println(&cfg)

	test := *cfg

	fmt.Println(test.App.Port)
	fmt.Println(test.Database)

}
