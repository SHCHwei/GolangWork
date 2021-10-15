package main

import (
	"fmt"
	"abc.com/flagT"
    "abc.com/dbcon"
)

func init(){
    dbcon.DBT()
}

func main(){

	fmt.Println("Hello world")
    flagT.GoflagT()

}