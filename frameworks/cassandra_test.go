package frameworks

import (
	"testing"
)

func TestCassandraBasic(t *testing.T) {
    println("awadwa")
	session := Init()
	defer session.Close()
	session.Query("select value from my table where pk1 = ?", "abc")
}

