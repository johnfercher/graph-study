package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func main() {
	qtd := 1000
	ids := []string{
		"f6043837-d32a-4a59-9c2f-ce8647a37a01",
		"ca1d2d28-93ae-4370-b1d1-50ded080a26b",
		"00f48479-7ce2-48ed-bf1a-f063eb8f7589",
		"5f916897-007a-4343-9047-bc1aaecb1a02",
		"35a5592c-b4ff-4178-be59-1598be97a250",
		"1f935058-9f66-4d31-967f-3e3c0a313602",
		"eb7974aa-738a-4edf-a555-f1e3c25db1ba",
		"2d00d5d0-eeda-4e80-a030-d001eff66ac7",
		"8e457f33-73c3-4ca3-8cb3-3ac577cb86e0",
		"c69454bc-de12-49e9-98b5-1259516ae473",
	}

	neo4, err := os.Create("test.cypher")
	if err != nil {
		fmt.Println(err)
		return
	}

	mysql, err := os.Create("test.sql")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range ids {
		for i := 0; i < qtd; i++ {
			newId, _ := uuid.NewRandom()

			_, errCreate := neo4.WriteString(fmt.Sprintf("CREATE (e:Vertex {id: '%s', type: 'E'});\n", newId))
			if errCreate != nil {
				fmt.Println(errCreate)
				neo4.Close()
				return
			}

			_, errAdd := neo4.WriteString(fmt.Sprintf("MATCH (d:Vertex), (e:Vertex) WHERE d.id = '%s' AND e.id = '%s' CREATE (d)-[:has]->(e);\n", id, newId))
			if errAdd != nil {
				fmt.Println(errAdd)
				neo4.Close()
				return
			}

			_, errCreate = mysql.WriteString(fmt.Sprintf("INSERT INTO vertex (`id`, `type`) VALUES ('%s', 'E');\n", newId))
			if errCreate != nil {
				fmt.Println(errCreate)
				mysql.Close()
				return
			}

			_, errAdd = mysql.WriteString(fmt.Sprintf("INSERT INTO edge (`parent_id`, `vertex_id`) VALUES ('%s', '%s');\n", id, newId))
			if errAdd != nil {
				fmt.Println(errAdd)
				mysql.Close()
				return
			}
		}
	}

	err = neo4.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
