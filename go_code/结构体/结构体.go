package main
import "fmt"

 //定义结构体
 type haha struct{ 
   
	name string
	age int
	color string

 }
 

 //多样的字段
 type lala struct{
    
	num1 [4]int
    num2 []int
	num3 map[string]string
	num4 *int

 }



func main(){

 

  var xixi haha = haha{"嘻嘻",19,"白"} //把结构体赋值给xixi


  xixi.name = "哈哈"
  xixi.age = 18
  xixi.color = "红"

  fmt.Println("xixi的信息为",xixi)
  fmt.Println("xixi的年龄为",xixi.age)
  fmt.Println("______________________________________")
  
  var haha lala
  haha.num1[0] = 10
  fmt.Println("字段为数组",haha.num1[0])

  //字段为切片类型时要make,才能使用
  haha.num2 = make([]int,5)
  haha.num2[0] = 100
  fmt.Println("字段为数组",haha.num1[0])
  
  //字段为map类型时也要make，才能使用
  haha.num3 = make(map[string]string)
  haha.num3["哈哈"] = "嘻嘻"
  fmt.Println("字段为map时",haha.num3["哈哈"])
   
  
  //字段为指针时,要给指针 指向的值
   var k = 10
   haha.num4 = &k
   fmt.Println("字段为指针类型时",*haha.num4)
   


//结构体是值类型
    // kk := haha
	// kk.num1[0] = 100
	// fmt.Println(kk.num1[0]) 
    // kk.num2 = make([]int,5)
	// kk.num2[0] = 10000 
	// fmt.Println(kk.num2[0])
	// fmt.Println("字段为数组",haha.num1[0])
	// fmt.Println("字段为数组",haha.num1[0])
	


}