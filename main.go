package main

import (
	"context"
	"fmt"
	"gopoc/cmd"
	"gopoc/internal/config"
	"gopoc/internal/db"
	"log"
	"path/filepath"
)

func main() {
	absPath, _ := filepath.Abs("ui/static")
	fmt.Println(absPath)
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
