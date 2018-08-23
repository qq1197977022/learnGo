package day04

import (
	"fmt"
	"reflect"
)

func CallMethod() {
	man := Man{"male", Human{"杨一帆", 22}}
	rValue := reflect.ValueOf(man)
	rType := reflect.TypeOf(man)
	for i := 0; i < rType.NumMethod(); i++ {
		m := rType.Method(i)
		fmt.Printf("%s\t %v\n", m.Name, m.Type)
	}
	rValue.Method(0).Call([]reflect.Value{reflect.ValueOf("苹果")})
	rValue.Method(1).Call([]reflect.Value{reflect.ValueOf(80)})
}
