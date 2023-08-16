package main

import(
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func checkerr(err error){
	if err != nil{
		panic(err)
	}
}

func fileread(filename string){
	data, err := ioutil.ReadFile(filename)
	checkerr(err)
	fmt.Println("Data in the file is\n", string(data))
}

func create(filename string){
	file, err := os.Create(filename)
	checkerr(err)
	defer file.Close()
	var content string
	fmt.Println("Enter data to be added to the file: ")
	fmt.Scanln(&content)
	trimmedcontent := strings.TrimSpace(content)
	_, err = file.Write([]byte(trimmedcontent))
	checkerr(err)
}

func main(){
	fmt.Println("File Handling in Golang")

	create("text.txt")

	fileread("./text.txt")

}