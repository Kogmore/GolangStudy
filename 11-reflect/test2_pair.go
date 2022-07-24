package main

import (
	"fmt"
	"io"
	"os"
)

/*
pair的介绍
*/

func main() {
	//tty:pair<type:*os.file , value=name 文件描述符>
	tty, err := os.OpenFile("C:/Windows/System32/WindowsPowerShell/v1.0/powershell.exe", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	//r:pair<type: , value:>
	var r io.Reader
	//r:pair<type:*os.file , value=name 文件描述符>
	r = tty
	//w:pair<type: , value:>
	var w io.Writer
	//w:pair<type:*os.file , value=name 文件描述符>
	w = r.(io.Writer)
	w.Write([]byte("HELLO THIS IS A TEST !!!!\n"))
}
