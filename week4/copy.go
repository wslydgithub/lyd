package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	type StringHeader struct {
	 Data uintptr
	 Len  int
	}

	type SliceHeader struct {
	 Data uintptr
	 Len  int
	 Cap  int
	}
*/

func Change(b []byte) string {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	newHeader := reflect.StringHeader{
		Data: header.Data,
		Len:  header.Len,
	}

	return *(*string)(unsafe.Pointer(&newHeader))
}

/*
&byte[0] 和 &byte 在 Go 语言中是不同的。
&byte[0] 表示获取字节切片 byte 的第一个元素的地址，即指向第一个元素的指针。
而 &byte 表示获取字节切片 byte 的地址，即指向整个字节切片的指针。
因此，它们是不同的指针类型，指向不同的内存位置。
*/
func main() {
	var bytes []byte
	bytes = []byte{'1', '2', '3', '4', '5'}
	fmt.Println(Change(bytes))
	fmt.Println(*(*string)(unsafe.Pointer(&bytes)))
	//var byte1 byte = 0
	//pb := *(*string)(unsafe.Pointer(&bytes))
	//pb2 := *(*string)(unsafe.Pointer(&byte1))
	//pb3 := pb + pb2
	//fmt.Println(pb3)

}
