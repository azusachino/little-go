package dialect

import (
	"github.com/azusachino/little-go/learn-project/orm/day_two/dialect"
	"reflect"
)

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSql(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect dialect.Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
