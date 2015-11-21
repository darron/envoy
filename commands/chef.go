package commands

import (
	"fmt"
	"github.com/go-chef/chef"
	"io/ioutil"
	"os"
)

func ReadKey(keypath string) string {
	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		Log(fmt.Sprintf("create: Could not read %s:", keypath), "info")
		os.Exit(1)
	}
	keyString := string(key)
	return keyString
}

func Connect(key, node, url string) *chef.Client {
	Log(fmt.Sprintf("create: Connecting to %s with '%s'", url, node), "info")
	client, err := chef.NewClient(&chef.Config{
		Name:    node,
		Key:     key,
		BaseURL: url,
		SkipSSL: true,
	})
	if err != nil {
		Log("create: Error with Chef connection", "info")
		os.Exit(1)
	}
	return client
}