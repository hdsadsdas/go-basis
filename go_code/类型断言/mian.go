package main
import "fmt"

//定义一个结构体

type haha struct{
    
	x int 
    y int

}

//定义一个空接口
//空接口才有类型断言

type xixi interface{

}


func main(){

   //定义一个空接口类型的变量
   var n1 xixi

   //定义一个结构体变量
   
   n2 := haha{20,10000}

   //将这个结构体赋给空接口
   n1 = n2

   //再定义一个结构体变量
   var n3 haha
   
   //n3 = n1 //空接口类型不能直接赋值给结构体变量
   n3 = n1.(haha)//类型断言  转换的类型要和接口接收的类型一样

   fmt.Println(n3)


   //给类型断言判断
   
   n4,ok := n1.(haha) //转换的类型要和接口接收的类型不时一样会报错
                    // 给一个 ok变量 只 接收对或错
					//这样不会报错 会继续执行下面的代码

	if ok {
        
         fmt.Println("类型断言成功")
		 fmt.Println(n4)
	}else{

		fmt.Println("类型断言失败")

	}				
   
    fmt.Println("代码继续执行")


}