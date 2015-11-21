package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "envoy",
	Short: "Create hosts file from Chef server nodes.",
	Long:  `Create hosts file from Chef server nodes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("`envoy -h` for help information.")
		fmt.Println("`envoy -v` for version information.")
	},
}

var (
	NodeName        string
	ChefServerUrl   string
	ChefEnvironment string
	ClientKey       string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&NodeName, "node", "n", "", "Node name with Chef server access.")
	RootCmd.PersistentFlags().StringVarP(&ChefServerUrl, "server", "s", "", "Chef Server url.")
	RootCmd.PersistentFlags().StringVarP(&ChefEnvironment, "environment", "e", "", "Chef Server environment.")
	RootCmd.PersistentFlags().StringVarP(&ClientKey, "key", "k", "", "Chef client key.")
}
