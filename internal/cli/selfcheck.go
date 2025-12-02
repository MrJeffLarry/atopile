package cli

import (
	"fmt"

	"github.com/atopile/atopile/internal/version"
	"github.com/spf13/cobra"
)

var selfCheckCmd = &cobra.Command{
	Use:    "self-check",
	Short:  "Quick self-check for extensions",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(selfCheckCmd)
}
