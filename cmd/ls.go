package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
	"github.com/spf13/cobra"
)

func NewLsCmd(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := tr.ListTasks()

			if len(tasks) == 0 {
				fmt.Println("You don't have any task yet, use the `tsk mk` command to make your first task!")
				return nil
			}

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			util.PrintTasks(tasks)

			return nil
		},
	}
}
