package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		taskService := task.NewTaskService()
		tasks, err := taskService.GetTasks()
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
