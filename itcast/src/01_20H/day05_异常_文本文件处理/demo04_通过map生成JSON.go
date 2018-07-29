package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	mp1 := make(map[string]interface{}, 4)
	mp1["Company"] = "itcast"
	mp1["Subjects"] = [4]string{"Go", "C++", "Python", "前端"}
	mp1["IsOk"] = true
	mp1["Price"] = 666

	mp2 := map[string]interface{}{
		"Company":  "itcast",
		"Subjects": [4]string{"Go", "C++", "Python", "前端"},
		"IsOk":     true,
		"Price":    666,
	}

	fmt.Printf("%+v, %d\n", mp1, len(mp1))

	fmt.Println("=========================================================")
	fmt.Printf("%+v, %d\n", mp2, len(mp2))

	jsnStr1, jsnStr2 := jsnDeal(mp1, mp2)
	fmt.Printf("%s\n-------------\n%s", jsnStr1, jsnStr2)
}
func jsnDeal(mp1, mp2 map[string]interface{}) (jsnStr1, jsnStr2 string) {
	bytSli, err := json.MarshalIndent(mp1, "", "		")
	if err != nil {
		log.Fatal(err)
	}
	jsnStr1 = bytes.NewBuffer(bytSli).String()

	bytSli, err = json.MarshalIndent(mp2, "", "		")
	if err != nil {
		log.Fatalln(err)
	}
	jsnStr2 = bytes.NewBuffer(bytSli).String()

	return
}
