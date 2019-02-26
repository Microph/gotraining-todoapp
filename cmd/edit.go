package cmd

import (
	"errors"
	"log"
	"os"
	"todoapp/todo"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit todo-name item-no new-item",
	Short: "Edit todo's item",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("requires todo-name, item-no and new-item")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Edit(args[0], args[1], args[2]); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}
