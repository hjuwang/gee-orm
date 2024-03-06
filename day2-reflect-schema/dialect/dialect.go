package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	// DataTypeOf 将go 类型 转化为数据库类型
	DataTypeOf(typ reflect.Value) string

	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
