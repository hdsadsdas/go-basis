package main

import (
	"fmt"
	"io/ioutil"
 )

func  main()  {
	
	n1 := "d:/wenjian"
	n2 := "d:/wenjian2"

	 file , ok := ioutil.ReadFile(n1)
			if ok != nil {
			
				fmt.Println("读取文件错误")

			}


    ok = ioutil.WriteFile(n2,file,0005)
	if ok != nil {

		fmt.Println("写入文件错误")

	}			 

}