package session

import (
	"fmt"
	"geeorm/dialect/schema"
	"geeorm/log"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	// nil or different model, update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {

	table := s.RefTable()

	var column []string

	for _, field := range table.Fields {
		column = append(column, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(column, ",")

	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()

	if err != nil {
		return err
	}

	return nil

}

func (s *Session) DropTable() error {

	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s;", s.RefTable().Name)).Exec()

	return err
}

func (s *Session) HashTable() bool {

	sql, value := s.dialect.TableExistSQL(s.refTable.Name)

	row := s.Raw(sql, value...).QueryRow()

	var tmp string

	_ = row.Scan(&tmp)

	return tmp == s.refTable.Name
}
