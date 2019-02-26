package cmd

import (
	"errors"
	"log"
	"os"
	"todoapp/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete todo-name item-no",
	Short: "Delete todo's item",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires todo-name and item-no")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Delete(args[0], args[1]); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}
