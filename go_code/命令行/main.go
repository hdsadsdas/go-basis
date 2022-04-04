package main

import(
	"fmt"
	_"os"
	"flag"
)

func main(){

	// fmt.Println("命令行参数有",len(os.Args))

	// for i , j := range os.Args{
	// 	fmt.Printf("args[%v]=%v\n",i,j)
	// } 


	//定义几个变量，用于接收命令行的参数值
	var n1 string
	var n2 string
	var n3 string
	var n4 int

	flag.StringVar(&n1,"u","","用户名，默认为空")
	flag.StringVar(&n2,"pwn","","密码，默认为空")
	flag.StringVar(&n3,"h","localhost","主机名，默认为localhost")
	flag.IntVar(&n4,"port",3306,"端口号，默认为3306")
	
	//转换，必须要调用
	flag.Parse()

   //输出结果

   fmt.Printf("用户名为:%v   密码为%v    主机名为%v     端口号为%v",n1,n2,n3,n4)

}