package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync local data into remote database",
	Run: func(cmd *cobra.Command, args []string) {
		taskService := task.NewTaskService()
		err := taskService.SyncTasks()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Done\n")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
