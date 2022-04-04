package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

type baochun struct{

	yingwen int 
    shuzi   int
	kongge  int
	qita    int


}

func main(){

    file ,ok := os.Open("d:/wenjian.txt")

	if ok != nil{
		fmt.Println("读取文件错误")
		return 
	}
   
	defer file.Close()

	var haha baochun

    read := bufio.NewReader(file)

	for {

		str , ok1 := read.ReadString('\n')
        if ok1 == io.EOF{
			break
		}
		//遍历得到的字符串

	

		for _ ,n := range str{

			switch  {
			case n >= 'a' && n <= 'z'  :
                      haha.yingwen++
			case n >= 'A' && n <= 'Z'  :
				      haha.yingwen++
			case n >= ' ' && n <= '\t'  :
				      haha.kongge++
			case n >= '0' && n <= '9'  :
				      haha.shuzi ++		  		  
            default :
			          haha.qita ++
				
			}


		}


	}

       fmt.Printf("英文为%v 数字为%v 空格为%v 其他为%v",haha.yingwen,haha.shuzi,haha.kongge,haha.qita)


}
	


