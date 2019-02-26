package cmd

import (
	"fmt"
	"os"
	"todoapp/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print list of todo list file (*.td)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.List(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
