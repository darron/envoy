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

func GetNodes(c *chef.Client) map[string]string {
	nodeList, err := c.Nodes.List()
	if err != nil {
		Log("create: Could not list nodes.", "info")
	}
	return nodeList
}

func GetNode(c *chef.Client, node string) chef.Node {
	nodeDetail, err := c.Nodes.Get(node)
	if err != nil {
		Log(fmt.Sprintf("create: could not get node info for '%s'", node), "info")
	}
	return nodeDetail
}
