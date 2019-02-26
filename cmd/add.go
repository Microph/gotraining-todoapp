package cmd

import (
	"errors"
	"log"
	"os"
	"todoapp/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add todo-name item",
	Short: "Add new todo's item",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires todo-name and item")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Add(args[0], args[1]); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}
