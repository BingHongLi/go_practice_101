package main

import (
	"fmt"
)

func main() {

	//mutable variabler declare
	shorthandVarable  := 123
	fmt.Println(shorthandVarable)

	// 顯性宣告型別
	var explicitDeclareTypeInt string = "i am string."
	fmt.Println(explicitDeclareTypeInt)

	// 一次宣告多個變數
	iAmBoolean, iAmString, iAmInt := false, "demo", 123
	fmt.Printf("this is boolean %t , this is string %s , this is int %d \n",iAmBoolean,iAmString,iAmInt)

	// 宣告常數變數
	const myStringHello string = "Hello, 世界"
	//myStringHello = ""  // will be error
	fmt.Println(myStringHello)

	// 在main function外宣告一個demoAdd function並呼叫
	fmt.Print(demoAdd(13,42))

}

/*

func 函數名(參數名 參數型別, 參數名 參數型別) 回傳內容型別 { 方法體 }

*/
func demoAdd(x int, y int) int {
	return x + y
}