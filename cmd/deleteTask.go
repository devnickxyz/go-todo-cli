package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"deletetask", "usun"},
	Short:   "Delete a task name",
	Long:    "Delete a task name - give one",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting the task with id: %s\n", args[0])
		DeleteTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
}

func DeleteTask(id string) error {
	sql := `
        DELETE FROM tasks where id = $1
    `
	commandTag, err := Pool.Exec(Ctx, sql, id)
	if err != nil {
		return fmt.Errorf("error deleting task %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		fmt.Printf("Cannot delete task with ID: %s\nAre you sure about it?", id)

	} else {
		fmt.Printf("Deleted tasks with ID: %s\n", id)
	}

	return nil
}
