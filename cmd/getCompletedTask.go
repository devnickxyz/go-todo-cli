package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCompletedTasks = &cobra.Command{
	Use:     "getcompleted",
	Aliases: []string{"completed", "sikalafon", "si"},
	Short:   "Shows all completd tasks",
	Long:    "Shows all completed task - long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing your done tasks")

		t, _ := GetCompletedTasks()
		PrintTasks(t)
	},
}

func init() {
	rootCmd.AddCommand(getCompletedTasks)
}

func GetCompletedTasks() ([]Task, error) {
	sql := `
		SELECT id, text, completed, created_at, updated_at
		FROM tasks
		WHERE completed = true
		ORDER BY created_at DESC
	`
	rows, err := Pool.Query(Ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("some trouble getting query done - see: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Text,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("opps, something wrong when scanning %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("xd cannot loop through all task - see: %w", err)
	}
	return tasks, nil
}

// func PrintTasks(tasks []Task) {
// 	if len(tasks) == 0 {
// 		fmt.Println("No tasks :(")
// 		return
// 	}
// 	for _, task := range tasks {
// 		status := "[ ]"
// 		if task.Completed {
// 			status = "[✓]"
// 		}
// 		// fmt.Println(k)
// 		fmt.Printf("%d. %s %s (Created: %s)\n",
// 			task.ID,
// 			status,
// 			task.Text,
// 			task.CreatedAt.Format("2006-01-02 15:04:05"))
// 	}
// }
