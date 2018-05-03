package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"os"
	"github.com/bitly/go-simplejson"
	"encoding/json"
)

func main() {

	// 把一個範例Json，寫入檔案中
	jsonFile, createErr := os.Create("demo.json")
	if(createErr != nil ){
		log.Println("can not create file.")
	}else{
		jsonFile.WriteString(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
			}
		}`)
	}

	// 讀檔
	fileData,readErr := ioutil.ReadFile("demo.json")
	if(readErr !=nil){
		log.Println("can not read file or file doesn't exists")
	}else{
		fmt.Println(string(fileData))
	}


	/*

		引用simplejson，進行json處理
		在import內加入
		"github.com/bitly/go-simplejson"
		在終端命令列內，使用dep ensure

	*/

	// 讀取檔案轉成可操作的json
	demoJsonFromFile,_ := simplejson.NewJson(fileData)

	// 取得欄位為字串的內容
	demoString,_ := demoJsonFromFile.Get("test").Get("string").String()
	fmt.Println(demoString)

	// 做類似傳統ORM
	type DemoORM struct {
		String_array []string
		Bool bool
		String string
		Int int
	}
	// 生成空實例
	var demoOrm DemoORM
	// 取得要轉換的byte
	demoJsonByte,_ := demoJsonFromFile.Get("test").EncodePretty()

	// 進行轉換，並把內容放回空實例中
	json.Unmarshal(demoJsonByte,&demoOrm)
	fmt.Println(demoOrm)

	// 將 instance 轉成 json
	ormToJson, _ :=json.Marshal(demoOrm)
	fmt.Println(string(ormToJson))

}
