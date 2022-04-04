package main
import "fmt"

//方法是作用在指定数据类型上的

type haha struct{
     
	name string
	age int
	 
}

type nana int  //给int个别名


//此方法只能通过 haha这个类型来调用 不能直接调用
//haha这个结构体有一个名为 lala 的方法

func (xixi haha) lala(){ //xixi为 变量名   haha 为绑定的类型  
	                     //lala为 方法名
   
	fmt.Println(xixi.name)

}

func (n *haha) yaya(){  //和haha指针类型绑定

	(*n).age = 18 //标准的使用方式
	// (*n).age = 18 也可以使用  n.age = 18  
	//  只有结构体才能这样使用  
    fmt.Println((*n).age)
}

func (n *nana) nana2(){  
     
	*n = *n + 2
  fmt.Println(*n)  

}



func main(){

  var kk haha //变量名 
   kk.name = "小狗"

   kk.lala() // 调用方法  kk传给xixi  xixi.name = kk.name 
             // 和函数传参类似
			 //只要类型为 haha 的 都有一个lala的方法 
   
	//传地址		 
   (&kk).yaya()
   //在go语言中  (&kk).yaya()等价于 kk.yaya() 
   //因为使用 方法 时编译器自动的给我们加上

    //  shu := 10 
    // var cheshi *int
    // cheshi = &shu
	// fmt.Println(cheshi)


   var n1 nana = 10 //nana是int的别名
    
   n1.nana2() // 调用方法




}