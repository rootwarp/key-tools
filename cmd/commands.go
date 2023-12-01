package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	root := cobra.Command{}
	root.AddCommand(getEthCommands())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
