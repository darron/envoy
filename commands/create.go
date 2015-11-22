package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a hosts file.",
	Long:  `create makes a hosts file from Chef Server data.`,
	Run:   startCreate,
}

func startCreate(cmd *cobra.Command, args []string) {
	checkFlags()

	chefKey := ReadKey(ClientKey)

	chefConn := Connect(chefKey, NodeName, ChefServerUrl)

	json := DoSearch(chefConn, ChefEnvironment)

	fmt.Println(json)
}

func checkFlags() {
	Log("create: Checking cli flags.", "debug")
	if FiletoWrite == "" {
		fmt.Println("Need a file to write hosts to. --file / -f")
		os.Exit(1)
	}
	if NodeName == "" {
		fmt.Println("Need a node name with access to the Chef Server. --node / -n")
		os.Exit(1)
	}
	if ChefServerUrl == "" {
		fmt.Println("Need a Chef Server URL. --server / -s")
		os.Exit(1)
	}
	if ChefEnvironment == "" {
		fmt.Println("Need a Chef environment. --environment / -e")
		os.Exit(1)
	}
	if ClientKey == "" {
		fmt.Println("Need a Chef client key. --key / -k")
		os.Exit(1)
	}
	Log("create: Required cli flags present.", "debug")
}

var (
	FiletoWrite string
)

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&FiletoWrite, "file", "f", "", "where to write the hosts file")
}
