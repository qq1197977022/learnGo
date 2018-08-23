package day04

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string "姓名"
	age  int    "年龄"
}

func GetTypeValue() {
	tom := Human{"汤姆", 21}
	typ := tom.getType()
	val := tom.getValue()
	fmt.Println(typ, val)
	fmt.Printf("tom\t type: %T\t Kind: %v\t value: %v\n", typ, typ.Kind(), val)
	fmt.Printf("包路径: %s\n", typ.PkgPath())
	fmt.Printf("类型名: %s\n", typ.Name())
}

func (h Human) getType() reflect.Type {
	return reflect.TypeOf(h)
}

func (h Human) getValue() reflect.Value {
	return reflect.ValueOf(h)
}
