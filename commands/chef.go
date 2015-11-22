package commands

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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

type SearchResult struct {
	Rows []struct {
		Data struct {
			Ipaddress string `json:"ipaddress"`
			Name      string `json:"name"`
		} `json:"data"`
		Url string `json:"url"`
	} `json:"Rows"`
	Start int `json:"Start"`
	Total int `json:"Total"`
}

type Node struct {
	Ipaddress string
	Name      string
}

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
