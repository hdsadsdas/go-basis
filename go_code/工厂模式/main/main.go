package main

import (
    "fmt"
    "go_code/工厂模式/haha"
)

func main(){
 
	n1 := haha.Xixi{"哈哈",18}
   
   fmt.Println(n1)//  值类型

//    n3 := n1
//    n3.Name = "asdad"
//    fmt.Println(n3)
//    fmt.Println(n1)


   n2 := haha.Gcms("嘻嘻",20) // 这是一个指针类型 
   fmt.Println(n2) // 打印出来的是地址
   
   fmt.Println(n2.Ff())
   //直接n2.name报错

   n2.Xg("啦啦")
   fmt.Println(n2.Ff())
}

