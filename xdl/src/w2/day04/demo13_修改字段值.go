package day04

import (
	"fmt"
	"reflect"
)

func ModifyFieldValue() {
	manPtr := &Man{"male", Human{"杨一帆", 21}}
	rValuePtr := reflect.ValueOf(manPtr)

	fmt.Println(manPtr, *manPtr, rValuePtr)

	value := rValuePtr.Elem().Field(0)
	if value.CanSet() {
		switch value.Kind() {
		case reflect.String:
			value.SetString("第三性别")
		case reflect.Int:
			value.SetInt(23)
		}
	}

	fmt.Println(manPtr, *manPtr, rValuePtr)
}
