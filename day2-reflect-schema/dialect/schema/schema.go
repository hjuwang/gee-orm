package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model  interface{}
	Name   string
	Fields []*Field

	FieldNames []string

	fieldMap map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse 将任意对象 解析为schema
func Parse(dest interface{}, d dialect.Dialect) *Schema {

	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {

		p := modelType.Field(i)

		//不是匿名字段，且可导出
		if !p.Anonymous && ast.IsExported(p.Name) {

			filed := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				filed.Tag = v
			}

			schema.Fields = append(schema.Fields, filed)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = filed

		}

	}

	return schema
}
