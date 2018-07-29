package demo01

import "fmt"

func ArrayLiteral() {
	arr1 := [3]int8{}
	arr2 := [3]string{"A", "B", "C"}
	arr3 := [...]bool{true, false, true}
	arr4 := [...]int{0: 1, 2: 2, 1: 3, 5: 5, 7}

	for i, v := range arr1 {
		fmt.Printf("arr1[%d] = %d\n", i, v)
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %s\n", i, arr2[i])
	}

	var i = 0
	for i < len(arr3) {
		fmt.Printf("arr3[%d] = %t\n", i, arr3[i])
		i++
	}
	for key, value := range arr4 {
		fmt.Println(key, value)
	}
}

func demo() {
	type Man struct {
		name string
		age  int
	}
	arr1 := [...]Man{{"梦航大神", 25}, {"包包", 25}}
	for i, v := range arr1 {
		fmt.Printf("%d\t%v", i, v)
	}

}
