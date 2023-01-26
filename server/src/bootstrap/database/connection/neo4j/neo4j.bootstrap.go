package neo4j

import (
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jConnection struct {
	Conn *neo4j.DriverWithContext
}

var neo4jSingleton *Neo4jConnection

func GetNeo4jConnection() *Neo4jConnection {
	if neo4jSingleton == nil {
		driver, err := neo4j.NewDriverWithContext(os.Getenv("NEO4J_HOST"), neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), ""))
		if err != nil {
			panic(err)
		}
		neo4jSingleton = &Neo4jConnection{Conn: &driver}
		log.Println("Neo4j connection created")
	}
	return neo4jSingleton
}
