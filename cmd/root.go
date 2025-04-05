package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

var Pool *pgxpool.Pool
var Ctx context.Context
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a cli tool for TODO app",
	Long:  "todo is great",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oop! Error when executing gotodo command %s\n", err)
		os.Exit(1)
	}
}
