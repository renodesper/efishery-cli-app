package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/models"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fID, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println(err)
			return
		}

		if fID == "" {
			fmt.Println("ID is required, please define the content with \"-i\"")
			return
		}

		fContent, err := cmd.Flags().GetString("content")
		if err != nil {
			fmt.Println(err)
			return
		}

		if fContent == "" {
			fmt.Println("Content is required, please define the content with \"-c\"")
			return
		}

		fTags, err := cmd.Flags().GetString("tags")
		if err != nil {
			fmt.Println(err)
			return
		}

		t := models.Task{
			ID:        fID,
			Content:   fContent,
			Tags:      fTags,
			CreatedAt: time.Now(),
		}

		taskService := task.NewTaskService()
		err = taskService.UpdateTask(&t)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Done\n")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("id", "i", "", "docID of the task")
	editCmd.Flags().StringP("content", "c", "", "Task content, wrap text with quotation mark")
	editCmd.Flags().StringP("tags", "t", "", "Tags for the task, separated by comma")
}
