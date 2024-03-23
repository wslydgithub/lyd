package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 自定义结构体
type Student struct {
	Number int
	Score1 [2]float64
	data   struct {
		Score map[string]float64
		Name  string
	}
}

// 处理不同类型
// 犯了个错误，不要用fmt.Sprintf/fmt.Sprintln
func jie(x interface{}, v reflect.Value) {
	var bytes []byte
	switch v.Kind() {
	case reflect.Invalid:
		bytes = []byte("nil")
		fmt.Printf("{\"%s\":\"%v\"}\n", x, bytes)
	case reflect.Int:
		fmt.Printf("{\"%s\":\"%d\"}\n", x, v.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("{\"%s\":\"%f\"}\n", x, v.Float())
	case reflect.String:
		fmt.Printf("{\"%s\":\"%s\"}\n", x, v.String())
	case reflect.Struct:
		fmt.Printf("{\"%s\":\n", x)
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i)
			fmt.Print("\t")
			jie(field.Name, value)
		}
		fmt.Printf("}\n")
	case reflect.Array, reflect.Slice:
		fmt.Printf("{\"%s\":\n", x)
		for i := 0; i < v.Len(); i++ {
			fmt.Print("\t")
			jie(strconv.Itoa(i), v.Index(i))
		}
		fmt.Printf("}\n")
	case reflect.Map:
		fmt.Printf("{\"%s\":\n", x)
		keys := v.MapKeys()
		for _, key := range keys {
			val := v.MapIndex(key)
			fmt.Print("\t")
			jie(key, val)
		}
		fmt.Printf("}\n")
	}
}

// 处理结构体
func Mar(x interface{}) {
	v := reflect.ValueOf(x)
	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)
		jie(field.Name, value)
	}

}

func main() {
	var student Student
	student.Score1[0] = 74.1
	student.Score1[1] = 90.1
	student.data.Name = "charlie"
	student.data.Score = make(map[string]float64)
	student.data.Score["高数"] = 74.1
	student.data.Score["C语言"] = 50
	student.Number = 2023
	Mar(&student)
}
