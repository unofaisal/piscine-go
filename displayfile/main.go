package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Print("File name missing\n")
	}
	if len(args) > 1 {
		fmt.Print("Too many arguments\n")
	}
	fName := args[0]
	file, err := os.Open(fName)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer file.Close()
	stat, _ := file.Stat()
	arr := make([]byte, stat.Size())
	file.Read(arr)
	fmt.Print(string(arr))
}
