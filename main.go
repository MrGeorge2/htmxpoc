package main

import (
	"context"
	"gopoc/cmd"
	"gopoc/internal/config"
	"gopoc/internal/db"
	"log"
)

func main() {
	ctx := context.Background()

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateOrGetDatabase(ctx)

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
