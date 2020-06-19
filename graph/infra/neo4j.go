package infra

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func GetNeo4jlConnection() (neo4j.Driver, error) {
	useConsoleLogger := func(level neo4j.LogLevel) func(config *neo4j.Config) {
		return func(config *neo4j.Config) {
			config.Log = neo4j.ConsoleLogger(level)
			config.Encrypted = false
		}
	}

	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "supersecret", ""), useConsoleLogger(neo4j.ERROR))
	if err != nil {
		return nil, err
	}

	return driver, nil
}
