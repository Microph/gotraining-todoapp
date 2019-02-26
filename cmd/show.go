package cmd

import (
	"errors"
	"log"
	"os"
	"todoapp/todo"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show todo-name",
	Short: "Show todo list of the todo-name",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires todo-name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Show(args[0]); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}
