package cmd

import (
	"context"
	"gopoc/internal/db"
	"gopoc/internal/db/repositories"
	"gopoc/internal/handlers"
	"log"

	"github.com/spf13/cobra"
)

var runApiCmd = &cobra.Command{
	Use:  "run-api",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runApi()
		if err != nil {
			log.Fatal(err.Error())
		}
		return err
	},
}

func runApi() error {
	ctx := context.Background()
	db, err := db.CreateOrGetDatabase(ctx)

	if err != nil {
		return err
	}

	return handlers.RegisterRouter(*repositories.New(db))
}

func init() {
	rootCmd.AddCommand(runApiCmd)
}
