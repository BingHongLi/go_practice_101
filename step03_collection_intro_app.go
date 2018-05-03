package main

import "fmt"

func main() {

	/*
		slice 不定長度的array
	*/
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	/*
		map demo
	*/
	mapDemo := map[string]int{}
	// set a pair
	mapDemo["key1"] = 999
	mapDemo["key2"] = 888
	mapDemo["key3"] = 777
	// get a pair
	fmt.Println(mapDemo["key1"])
	// delete a pair
	delete(mapDemo, "key1")
	// list all pair in map
	fmt.Println(mapDemo)


	/*
		struct demo
		資料集合之定義
	*/
	type person struct {
		name string
		age  int
	}
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	// 不一定要像類別一樣，把建構子填滿
	fmt.Println(person{name: "Fred"})

	demoPerson := person{name: "Sean", age: 50}
	fmt.Println(demoPerson.name)

}

