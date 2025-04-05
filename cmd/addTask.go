package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addtask", "dodaj"},
	Short:   "Add new task name",
	Long:    "Add new new task name - give one",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding new tasks %s\n", args[0])
		CreateTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}

func CreateTask(text string) error {
	sql := `
        INSERT INTO tasks (text, completed)
        VALUES ($1, $2)
        RETURNING id
    `
	var id int
	var err error
	err = Pool.QueryRow(Ctx, sql, text, false).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating task %w", err)
	}
	fmt.Printf("Created task with ID: %d\n", id)
	return nil
}
