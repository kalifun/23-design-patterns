package builder

import "fmt"

// 建立一个db对象
type DB struct {
	Table  interface{}
	Filter []string
}

// 先创建一个db对象
func NewOrm() *DB {
	return &DB{}
}

// table 相关
func (d *DB) QueryTable(i interface{}) *DB {
	d.Table = i
	return d
}

// where 语句
func (d *DB) Where(field string, v interface{}) *DB {
	s := fmt.Sprintf("%s %v", field, v)
	d.Filter = append(d.Filter, s)
	return d
}

// 查询all
func (d *DB) All() *DB {
	return d
}
