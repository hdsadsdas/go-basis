 package main

import (
	"fmt"
	"os" // 一个控制打开文件的包
	"io"
	"bufio"
)

func main(){
                                         //./main.go   相对路径
	dakai,err := os.Open("d:/wenjian.txt") //os包里面的一个函数用来打开文件
	                                //输入文件地址  返回的是一个 文件（file）结构体指针
 
	if err != nil {
		fmt.Println("打开文件失败")
	}
 
	fmt.Printf("打开的文件的地址%v\n",dakai)
     
	//defer 当函数退出时才会操作
      defer dakai.Close()//os包里的一个方法
                         //关闭文件
     

		//这里会返回一个指针	带缓冲的
		//文件较大 使用这个			 
		reader := bufio.NewReader(dakai)				 

  for {
	  str , err := reader.ReadString('\n') //str 每一行的信息
	                                       //err 判断有没有读到错误
    
		 fmt.Print(str) //这里可以不用换行
		             //输出每一行的信息
	           // 因为上面的"/n"会加到str里										   

	  if err == io.EOF{ //io.EOF也是一个错误 指到尾了
		                //当文件读取带末尾结束
		  break
	  }
	  
  }
  fmt.Println("文件读取结束")


}