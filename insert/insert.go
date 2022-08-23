package insert

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

func InsertStmt(entity interface{}) (string, []interface{}, error) {
	if entity == nil {
		return "", nil, errInvalidEntity
	}
	val := reflect.ValueOf(entity)
	typ := val.Type()

	// 检测 entity 是否符合我们的要求
	// 我们只支持有限的几种输入
	// 不支持多级指针
	if typ.Kind() == reflect.Ptr {
		if typ.Elem().Kind() == reflect.Ptr {
			return "", nil, errInvalidEntity
		}
		typ = typ.Elem()
		val = val.Elem()
	}
	// 不支持空结构体
	if typ.Kind() == reflect.Struct && typ.NumField() == 0 {
		return "", nil, errInvalidEntity
	}

	num := typ.NumField()
	bd := strings.Builder{}
	bd.WriteString("INSERT INTO ")
	bd.WriteString("`" + typ.Name() + "`" + "(")
	//args := make([]interface{}, num)
	for i := 0; i < num; i++ {
		refVal := val.Field(i).Interface()
		fmt.Println("val: ", i, refVal)
		//args = append(args, val.Field(1).Interface())
		fd := typ.Field(i)
		bd.WriteString("`" + fd.Name + "`")
		if i == num-1 {
			bd.WriteString(") ")
		} else {
			bd.WriteString(",")
		}
	}
	bd.WriteString("VALUES(")
	for i := 0; i < num; i++ {
		bd.WriteString("?")
		if i == num-1 {
			bd.WriteString(");")
		} else {
			bd.WriteString(",")
		}
	}
	fmt.Println(bd.String())
	return bd.String(), nil, nil
	// 使用 strings.Builder 来拼接 字符串
	// bd := strings.Builder{}

	// 构造 INSERT INTO XXX，XXX 是你的表名，这里我们直接用结构体名字

	// 遍历所有的字段，构造出来的是 INSERT INTO XXX(col1, col2, col3)
	// 在这个遍历的过程中，你就可以把参数构造出来
	// 如果你打算支持组合，那么这里你要深入解析每一个组合的结构体
	// 并且层层深入进去

	// 拼接 VALUES，达成 INSERT INTO XXX(col1, col2, col3) VALUES

	// 再一次遍历所有的字段，要拼接成 INSERT INTO XXX(col1, col2, col3) VALUES(?,?,?)
	// 注意，在第一次遍历的时候我们就已经拿到了参数的值，所以这里就是简单拼接 ?,?,?

	// return bd.String(), args, nil
	panic("implement me")
}
