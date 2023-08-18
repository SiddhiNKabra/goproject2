package main

import(
	"fmt"
	"os"
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
	content := "Welcome to Golang.\n This is file hanling in golang.\n golang is a very versatile language.\n golang is Google's official language."
	_, err = file.WriteString(content)
	checkerr(err)
}

func main(){
	fmt.Println("File Handling in Golang")

	create("text.txt")

	fileread("./text.txt")

}