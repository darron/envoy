package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootCmd is the base command that sets up all other commands.
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
	// NodeName is the Chef node with access to the Chef server API.
	NodeName string

	// ChefServerURL is the URL for the Chef server.
	ChefServerURL string

	// ChefEnvironment is the Chef environment you are querying.
	ChefEnvironment string

	// ClientKey is the Chef client key that pairs with the NodeName.
	ClientKey string

	// DogStatsd sends stats to a Dogstatsd daemon if true.
	DogStatsd bool

	// DogStatsdAddress is the address of the Dogstatsd daemon.
	DogStatsdAddress string

	// DatadogAPIKey is the API key for the Datadog API.
	DatadogAPIKey string

	// DatadogAPPKey is the App key for the Datadog API.
	DatadogAPPKey string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&NodeName, "node", "n", "", "Node name with Chef server access.")
	RootCmd.PersistentFlags().StringVarP(&ChefServerURL, "server", "s", "", "Chef Server url.")
	RootCmd.PersistentFlags().StringVarP(&ChefEnvironment, "environment", "e", "", "Chef Server environment.")
	RootCmd.PersistentFlags().StringVarP(&ClientKey, "key", "k", "", "Chef client key.")
	RootCmd.PersistentFlags().BoolVarP(&DogStatsd, "dogstatsd", "d", false, "Send metrics to Dogstatsd")
	RootCmd.PersistentFlags().StringVarP(&DogStatsdAddress, "dogstatsd_address", "D", "localhost", "Address for dogstatsd server.")
	RootCmd.PersistentFlags().StringVarP(&DatadogAPIKey, "datadog_api_key", "a", "", "Datadog API Key")
	RootCmd.PersistentFlags().StringVarP(&DatadogAPPKey, "datadog_app_key", "A", "", "Datadog App Key")
}
