package main

import (
	"fmt"
	"io/ioutil"
)

func main(){

	haha := "d:/wenjian.txt"
   
	 a,ok :=ioutil.ReadFile(haha) //直接一次性打开   只适用于打开小文件
                        // 这个函数直接将打开关闭弄一起了所有不用再关闭

		if ok != nil {
            fmt.Println("读取文件有误")
		}				

						//这里转型的 原因是 上面函数返回的是[]byte切片
     fmt.Printf("%v",string(a))

}