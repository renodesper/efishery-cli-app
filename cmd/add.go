package cmd

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/renodesper/efishery-cli-app/models"
	"gitlab.com/renodesper/efishery-cli-app/services/task"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
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

		id, err := uuid.NewUUID()
		if err != nil {
			fmt.Println(err)
			return
		}

		t := models.Task{
			ID:        id.String(),
			Content:   fContent,
			Tags:      fTags,
			CreatedAt: time.Now(),
		}

		taskService := task.NewTaskService()
		err = taskService.AddTask(&t)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Done\n")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("content", "c", "", "Task content, wrap text with quotation mark")
	addCmd.Flags().StringP("tags", "t", "", "Tags for the task, separated by comma")
}
