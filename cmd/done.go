package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fID, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println(err)
			return
		}

		taskService := task.NewTaskService()
		err = taskService.DoneTask(fID)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Done\n")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.Flags().StringP("id", "i", "", "docID of the task")
}
