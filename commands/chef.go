package commands

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chef/chef"
	"io/ioutil"
	"os"
	"time"
)

// SearchResult is the struct that describes the Chef server search result.
type SearchResult struct {
	Rows []struct {
		Data struct {
			Ipaddress string `json:"ipaddress"`
			Name      string `json:"name"`
		} `json:"data"`
		URL string `json:"url"`
	} `json:"Rows"`
	Start int `json:"Start"`
	Total int `json:"Total"`
}

// Node decscribes the data that we actually want from the SearchResult.
type Node struct {
	Ipaddress string
	Name      string
}

// ReadKey reads the PEM encoded key for access to the Chef server.
func ReadKey(keypath string) string {
	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		Log(fmt.Sprintf("create: Could not read %s:", keypath), "info")
		os.Exit(1)
	}
	keyString := string(key)
	return keyString
}

// Connect sets up a connection to the Chef server and passed back a client.
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

// GetNodes returns a list of nodes that are present in a particular
// environment on a Chef server.
func GetNodes(c *chef.Client, env string) []Node {
	jsonString := DoSearch(c, env)
	jsonBytes := []byte(jsonString)
	nodes := CleanSearchResult(jsonBytes)
	return nodes
}

// DoSearch uses the Chef API to search for all nodes in a particular
// Chef server.
func DoSearch(c *chef.Client, env string) string {
	part := make(map[string]interface{})
	part["name"] = []string{"name"}
	part["ipaddress"] = []string{"ipaddress"}
	search := fmt.Sprintf("chef_environment:%s AND *:*", env)
	pres, err := c.Search.PartialExec("node", search, part)
	if err != nil {
		Log("create: Error with Chef partial search.", "info")
	}
	jsonData, _ := json.MarshalIndent(pres, "", "\t")
	jsonString := string(jsonData)
	return jsonString
}

// CleanSearchResult takes the JSON string from the Chef server search and
// only returns the Nodes that are present.
func CleanSearchResult(jsonBytes []byte) []Node {
	nodes := []Node{}
	var data SearchResult
	err := json.Unmarshal(jsonBytes, &data)

	if err != nil {
		spew.Dump(err)
	}

	for _, node := range data.Rows {
		ip := node.Data.Ipaddress
		name := node.Data.Name
		newNode := Node{Ipaddress: ip, Name: name}
		nodes = append(nodes, newNode)
	}
	return nodes
}

// RenderFile takes a list of Nodes and concatentates a string for the
// Nodes that have both a name AND and ip address.
func RenderFile(n []Node) string {
	t := time.Now()
	text := fmt.Sprintf("# Built on %s\n", t.UTC().Format(time.UnixDate))
	text += fmt.Sprintf("# For environment '%s' on '%s'\n", ChefEnvironment, ChefServerURL)
	for _, node := range n {
		if node.Name != "" && node.Ipaddress != "" {
			each := fmt.Sprintf("%s %s\n", node.Ipaddress, node.Name)
			text += each
		}
	}
	return text
}
