package main

import (
	"fmt"
	"github/test2"
	 "github/json_read"
	  "github/folder1"
)

func main() {
	fmt.Print("HI HOW ARE YOU")

	var rectLen, rectWidth float64 = 6, 7
	fmt.Printf("area of rectangle %.2f\n", test2.Area(rectLen, rectWidth));

	fmt.Println(test2.Diagonal(rectLen,rectWidth));

       json_read.Json_File_Read()
	folder1.Json_File_Read1()

}
