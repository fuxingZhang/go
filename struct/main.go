package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" form:"username"`
	Age  int    `json:"age" form:"age"`
}

func (p User) Get(prop string) reflect.Value {
	v := reflect.ValueOf(p)
	return v.FieldByName(prop)
}

func (p *User) Set(prop string, val string) bool {
	v := reflect.ValueOf(p)
	field := v.Elem().FieldByName(prop)
	fmt.Printf("%v,%T", field, field)
	field.SetString(val)
	return true
}

func main() {
	user := User{
		Name: "zfx",
		Age:  30,
	}
	fmt.Println(user.Get("Name"))
	fmt.Println(user.Set("Name", "zfx_updated"))
	fmt.Println(user.Get("Name"))
	fmt.Println(user)
}
