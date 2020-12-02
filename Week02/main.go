package main

import (
	"Week02/biz"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(biz.AddData(1, "wangrx"))
	fmt.Println(biz.AddData(0, "wangrx"))
	fmt.Println(biz.AddData(2, ""))
}
