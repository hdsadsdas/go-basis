package main

import "fmt" 

type usb interface{ //创建接口
	                //接口也有继承性 用了继承要实现继承过来接口的方法
					//interface 是引用类型
    
	haha()  //方法
    xixi()

}

//创建结构体
type n1 struct{

}
//创建结构体
type n2 struct{

}

//实现usb的方法 //实现了接口
func (n n1) haha(){

	fmt.Println("n1哈哈")

}

func (n n1) xixi(){

	fmt.Println("n1嘻嘻")

}

func (n n2) haha(){

	fmt.Println("n2哈哈")

}

func (n n2) xixi(){

	fmt.Println("n2嘻嘻")

}

func jk(usb usb){


}


//创建一个接口对应的
type computer struct{

}

//编写一个方法
func (c computer) chuan(n usb){

	n.haha()
	n.xixi()
    

}

//定义一个空接口
//我们可以把任何一个变量赋给空接口
type kong interface{

}




func main(){

	//测试
	c := computer{}  //创建一个computer结构体
	
	n1 :=n1{}//创建一个n1结构体
    n2 :=n2{}//创建一个n2结构体
   
	c.chuan(n1)
    c.chuan(n2)

	var jj usb = n1 //用接口来接收一个结构体
   jj.haha()//调用接口里面的方法
   
     

}