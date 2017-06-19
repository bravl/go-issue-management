package main

import (
	"fmt"
	"os"
	//    "time"
	"bytes"
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Issue  issueInfo
	Server serverInfo
}

type issueInfo struct {
	IssueNumber string
	Description string
	IssueLink   string
}

type serverInfo struct {
	Ip       string
	User     string
	Password string
}

func retrieveParameters() (tomlConfig, int) {
	var config tomlConfig

	if _, err := toml.DecodeFile(".conf", &config); err != nil {
		fmt.Println(err)
		return config, -1
	}
	return config, 0

}

func createIssue(nr string, desc string) {
	var config tomlConfig
	config.Issue.IssueNumber = nr
	config.Issue.Description = desc
	config.Issue.IssueLink = "https://itrack.barco.com/browse/" + nr
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		return
	}
	fmt.Println(buf.String())
}

func main() {
	var config tomlConfig
	var err int
	argsWithProg := os.Args[1:]
	fmt.Println(argsWithProg)

	if config, err = retrieveParameters(); err != 0 {
		fmt.Println("Failed to retrieve parameters")
		return
	}
	fmt.Printf("Issue: %s\n", config.Issue.IssueNumber)
	fmt.Printf("Description: %s\n", config.Issue.Description)
	fmt.Printf("Server ip: %s\n", config.Server.Ip)
	createIssue("CS0040-247", "Test description")
}
