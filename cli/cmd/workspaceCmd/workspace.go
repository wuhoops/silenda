/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package workspaceCmd

import (
	"github.com/wuhoops/silenda/cli/cmd"

	"github.com/spf13/cobra"
)

// workspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	cmd.RootCmd.AddCommand(workspaceCmd)
}
