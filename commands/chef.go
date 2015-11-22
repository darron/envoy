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

func GetNodes(c *chef.Client, env string) []Node {
	jsonString := DoSearch(c, env)
	jsonBytes := []byte(jsonString)
	nodes := CleanSearchResult(jsonBytes)
	return nodes
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

func RenderFile(n []Node) string {
	t := time.Now()
	text := fmt.Sprintf("# Built on %s\n", t.UTC().Format(time.UnixDate))
	text += fmt.Sprintf("# For environment '%s' on '%s'\n", ChefEnvironment, ChefServerUrl)
	for _, node := range n {
		each := fmt.Sprintf("%s %s\n", node.Name, node.Ipaddress)
		text += each
	}
	return text
}
