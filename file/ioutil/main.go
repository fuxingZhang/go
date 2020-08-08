package main

import (
	"fmt"
	"io/ioutil"
)

func read() {
	byteStr, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}

func write() {
	str := "ioutil"
	err := ioutil.WriteFile("../write.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	read()
	write()
}
