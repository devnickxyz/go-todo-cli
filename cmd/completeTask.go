package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeTaskCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"makedone", "done"},
	Short:   "Complete a task",
	Long:    "Complete a task name - give one",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Completing the task with id: %s\n", args[0])
		CompleteTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeTaskCmd)
}

func CompleteTask(id string) error {
	sql := `
        UPDATE tasks
		SET completed = true, updated_at = NOW()
		WHERE id = $1
    `
	commandTag, err := Pool.Exec(Ctx, sql, id)
	if err != nil {
		return fmt.Errorf("error completing task %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		fmt.Printf("Cannot complete task with ID: %s\nAre you sure about it?", id)

	} else {
		fmt.Printf("Completed tasks with ID: %s\n", id)
	}

	return nil
}
