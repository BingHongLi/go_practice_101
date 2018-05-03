package main

import (
	"fmt"
	"os"
	"io"
)

/*

controll flow demo

*/
func main() {

	/*
		if 條件 {方法體}
		else if 條件 {方法體}
		else {方法體}
	*/
	num := 10
	if  num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/*
		switch demo
	*/
	switchString := "linux"
	switch switchString {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", switchString)
	}

	/*
		for loop demo
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}

/*
	defer demo
	一般應用場景，用在把資源關閉的時候。
*/
func CopyFile(dstName, srcName string) (written int64, err error) {
	/*
		開啟檔案，CopyFile執行完成後，會執行src.Close() 與 dst.Close()
	*/
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}