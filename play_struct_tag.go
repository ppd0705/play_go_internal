package main

import (
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string `key1:"value1" key2:"value2"`
	ServerIP string `key3:"world"`
}

func main() {
	s := Server{"Hello", "World"}
	st := reflect.TypeOf(s)
	field1 := st.Field(0)
	field2, _ := st.FieldByName("ServerIP")
	fmt.Printf("key1 %s\n", field1.Tag.Get("key1"))
	fmt.Printf("unknown %s\n", field1.Tag.Get("key100"))
	fmt.Printf("key3 %s\n", field2.Tag.Get("key3"))
}