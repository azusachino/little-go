package frameworks

import "github.com/gocql/gocql"

// Init, the caller should handle resource leak
func Init() *gocql.Session {
	cluster := gocql.NewCluster("172.0.0.1:7000")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Any
	cluster.ProtoVersion = 4

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	return session
}
