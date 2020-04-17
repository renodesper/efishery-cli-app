package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks (offline-first, call remote database when connection is available)",
	Run: func(cmd *cobra.Command, args []string) {
		taskService := task.NewTaskService()
		tasks, outdatedTasks, err := taskService.GetTasks()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(fmt.Sprintf("docID\t\t\t\t\tStatus\tTags\tContent"))
		for _, t := range tasks {
			status := "Todo"

			if t.IsDone {
				status = "Done"
			}

			fmt.Println(fmt.Sprintf("%s\t%s\t%s\t%s", t.ID, status, t.Tags, t.Content))
		}

		fmt.Println(fmt.Sprintf("\nOutdated tasks: %d", len(outdatedTasks)))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
