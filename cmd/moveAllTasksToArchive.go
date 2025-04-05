package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var moveAllTaskToArchive = &cobra.Command{
	Use:     "archiveall",
	Aliases: []string{"arch"},
	Short:   "Archive",
	Long:    "Archive all taks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Archiving your all tasks")
		MoveAllTasksToArchive()
	},
}

func init() {
	rootCmd.AddCommand(moveAllTaskToArchive)
}

func MoveAllTasksToArchive() error {
	tx, err := Pool.Begin(Ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w\n", err)
	}
	defer tx.Rollback(Ctx)

	sql := `
		INSERT INTO tasks_archive (id, text, completed, created_at, updated_at)
		SELECT id, text, completed, created_at, updated_at FROM tasks
	`
	_, err = tx.Exec(Ctx, sql)

	if err != nil {
		return fmt.Errorf("error coping to archive: %w\n", err)
	}
	if err = tx.Commit(Ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w\n", err)
	}

	return nil
}
