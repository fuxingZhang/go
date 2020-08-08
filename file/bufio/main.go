package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read() {
	file, err := os.Open("../test.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(file) // &{0xc00007e780}
	reader := bufio.NewReader(file)
	var fileStr string
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fileStr += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)
}

func write() {
	file, err := os.OpenFile("../write.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("11111\r\n")
	writer.Write([]byte("22222\r\n"))
	writer.Flush()
}

func main() {
	read()
	write()
}
