package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type T2 struct {
	n1 int
	N2 int
}
type T3 struct {
	n1 int
	N2 int `json:"N1"`
}

func main() {
	var emin interface{}
	flo := 5.111
	btslice, _ := json.Marshal(flo)
	t3 := T3{} //根据`json`进行识别

	_ = json.Unmarshal(btslice, &emin)
	_ = json.Unmarshal(btslice, &t3)
	fmt.Println("emin:", emin)
	fmt.Println("type1", reflect.TypeOf(emin))
	fmt.Println("btslice:", string(btslice))
	fmt.Println(t3)
	//进行json解析时，若以interface{}接收数据，则会按照下列规则进行解析：
	//
	//bool, for JSON booleans
	//float64, for JSON numbers
	//string, for JSON strings
	//[]interface{}, for JSON arrays
	//map[string]interface{}, for JSON objects
	//nil for JSON null

	t2 := T2{1, 2}
	btsli, _ := json.Marshal(t2)
	_ = json.Unmarshal(btsli, &emin)
	fmt.Println("type2", reflect.TypeOf(emin))
	fmt.Println(emin)

}
