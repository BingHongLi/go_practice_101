package main

import (
	"os"
	"fmt"
	"errors"
)


func main() {

	/*
		接受Error，並進行處理
	*/
	f, err := os.Open("filename.ext")
	if err != nil {
		fmt.Println("檔案不存在")
	}else{
		f.Close()
	}

	/*
		接收客製化function的error
	*/
	result, errResult := genError(5)
	if(errResult==nil){
		fmt.Println(result)
	}else{
		fmt.Println(result)
	}

}

/*
	客製化生成Error
*/
func genError(needMoreThan10 int) (string, error) {
	if needMoreThan10 > 10 {
		return "成功", nil
	}else{
		return "失敗", errors.New("變數小於10,出錯")
	}

}