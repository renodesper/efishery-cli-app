package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fID, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println(err)
			return
		}

		taskService := task.NewTaskService()
		err = taskService.DeleteTask(fID)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Done\n")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("id", "i", "", "docID of the task")
}
