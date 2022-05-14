package cmd

import (
	"KDRC-Client/cmd/account"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kdrc",
	Short: "KangDroid-Cloud Client Command",
	Long:  "Main Entry point!",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(account.AccountCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
