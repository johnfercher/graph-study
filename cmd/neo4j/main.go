package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	response, err := helloWorld("bolt://localhost:7687", "neo4j", "supersecret")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func helloWorld(uri, username, password string) (interface{}, error) {
	var (
		err      error
		driver   neo4j.Driver
		session  neo4j.Session
		result   neo4j.Result
		greeting interface{}
	)

	useConsoleLogger := func(level neo4j.LogLevel) func(config *neo4j.Config) {
		return func(config *neo4j.Config) {
			config.Log = neo4j.ConsoleLogger(level)
			config.Encrypted = false
		}
	}

	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), useConsoleLogger(neo4j.ERROR))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	session, err = driver.Session(neo4j.AccessModeRead)
	if err != nil {
		return "", err
	}
	defer session.Close()

	greeting, err = session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err = transaction.Run(
			"MATCH (v:Vertex)-[r]->() WHERE v.id = $id RETURN COUNT(r)",
			map[string]interface{}{"id": "c37e04a4-01e6-4b6f-bb34-94bc60dd1495"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting, nil
}
