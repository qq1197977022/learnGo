package day04

import (
	"fmt"
	"reflect"
)

type Man struct {
	Gender string
	Human
}

func GetFieldInfo() {
	tom := Man{"male", Human{"汤姆喵", 12}}
	tom.fieldInfo()
}

func (m Man) fieldInfo() {
	typ := reflect.TypeOf(m)
	fmt.Println(typ)
	for i := 0; i < typ.NumField(); i++ {
		fmt.Printf("%+-10v\n", typ.Field(i))

		fmt.Println("------------------------------------------------------")
		fmt.Println(typ.FieldByIndex([]int{1, 0}), typ.FieldByIndex([]int{1, 1}))
	}
}
