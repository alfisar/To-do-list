package main

import (
	"context"
	"log"
	"todolist/config"
	"todolist/router"
)

func main() {

	ctx := context.Background()
	_, err := config.InitConfig(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	route := router.NewRouter()
	route.Listen(":8080")
}
