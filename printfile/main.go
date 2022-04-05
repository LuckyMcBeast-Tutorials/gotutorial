package main

import (
	"fmt"
	"io"
	"os"
)

var filepath string = os.Args[1]

func main(){
	_, statErr := os.Stat(filepath)
	handleStatErr(statErr)

	//Option 1 - original choice
	bs, readErr := os.ReadFile(filepath)
	handleOpenErr(readErr)
	fmt.Println(string(bs))

	//Option 2 - instructor provided
	op, openErr := os.Open(filepath)
	handleOpenErr(openErr)
	io.Copy(os.Stdout, op)
	
	fmt.Println()
}

func handleStatErr(e error){
	if e != nil {
		os.Create(filepath)
	}
}

func handleOpenErr(e error){
	if e != nil {
		fmt.Errorf("Failed to open file named %v", filepath)
	}
}