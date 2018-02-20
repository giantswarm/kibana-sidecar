package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCmd just prints something.
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Make architect happy by having a version command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This command is actually not implemented.")
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
