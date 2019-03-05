package cmd

import (
	"errors"
	"log"
	"os"
	"todoapp/todo"

	"github.com/spf13/cobra"
)

var pyramidCmd = &cobra.Command{
	Use:   "pyramid number",
	Short: "Print pyramid",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires number input")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Pyramid(args[0]); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}
