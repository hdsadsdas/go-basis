package main

import "fmt"

type haha struct{

	name string
	age int

}

type ym struct{
	name string
}

func (n *haha) haha1(){
 
	fmt.Println("name=",n.name)

}

type xixi struct{
    ming ym //有名结构体  如果是一个有名结构体使用时就必须带上 她的名字
	haha  // 继承haha的所有方法和字段 匿名字段
    red int //也可以有自己独有的字段和方法
	name string // 当与匿名结构体 重复 要用 n2.haha.name 和 n2.name 区分
}

func (n *xixi) xixi1(){
 
	fmt.Println("name=",n.red)

}

type jiji struct{

	haha 
	red int 

}

func (n *jiji) jiji1(){
 
	fmt.Println("name=",n.red)

}

type kaka struct{

	*haha  //用指针类型提高效率   但要传入地址
	

}


func main(){
 
fmt.Println("n1")
	var n1 haha
     n1.name = "哈哈"	
     n1.age = 19
	 n1.haha1()

 fmt.Println("n2")
	var n2 xixi
	 n2.haha.name = "嘻嘻" // 可简写为n2.name
	 n2.haha.age = 19
	 n2.red = 5000
	 n2.haha.haha1()
     n2.xixi1()

 fmt.Println("n3")
	 var n3 jiji
	 n3.haha.name ="叽叽"
	 n3.haha.age = 20
	 n3.red = 10000
     n3.haha.haha1()
	 n3.jiji1()

fmt.Println("n4")
	 //嵌套结构体的直接赋值
	  n4 := jiji{haha{"直接赋值",21},15000}
	  n4.haha.haha1()
	  n4.jiji1()

fmt.Println("n5")
	  
	   n5 := kaka{&haha{"指针嵌套",21}}
	   n5.haha.haha1()
	   	  

}