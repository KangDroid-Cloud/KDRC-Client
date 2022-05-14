package account

import (
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account Related Commands",
	Long:  "Get account information, or register to service.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AccountCmd.AddCommand(registrationCommand)
}
