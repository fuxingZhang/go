package main

import (
	"fmt"
	"io"
	"os"
)

func read() {
	file, err := os.Open("./test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	var tmpSlice = make([]byte, 128)
	var strSlice []byte
	for {
		n, err := file.Read(tmpSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		strSlice = append(strSlice, tmpSlice[:n]...)
		fmt.Printf("read bytesLength: %v\n", n)
	}
	fmt.Println(string(strSlice))
}

func write() {
	file, err := os.OpenFile("./write.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("11111\r\n")
	file.Write([]byte("22222\r\n"))
}

func main() {
	read()
	write()
}
