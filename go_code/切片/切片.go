package main
import "fmt"

func main(){
	//第一种方式
	var haha =[6]int {1,2,3,4,5,6}
	    xixi := haha[1:2]//不包含下标2   
	fmt.Println("第一种方式",xixi)
 

	//第二种方式
	var n1 []int  = make([]int ,5 ,10)//make内置函数  需要三个参数 （类型  长度  容量）
     n1[1]=2
	 n1[2]=2
   fmt.Println("第二种方式",n1)

    
   //第三种方式
    var n2 = []int {1,2,3,4,5}
	fmt.Println("第三种方式",n2) 

   //切片的添加
   n2 = append(n2,10,20,30)
   fmt.Println("添加1",n2)
                       
   n1 = append(n1,n1...)//可以添加切片
   fmt.Println("添加2",n1)
   
   xixi = append(xixi,1,2,3,4,42,64,7)
   fmt.Println("添加3",xixi)
   xixi[0] = 100
   fmt.Println("引用的数组",haha)//被引用的数组 会断开连接 不再引用
   
   //copy
   var n3 = []int{1,2,3,4}
   var n4 []int = make([]int,1)//长度只有1 所以拷贝过来只有一个

   copy(n4,n3)//n4 和 n3 没有联系

   fmt.Println("拷贝",n4)

   //string 和 切片的关系

     n5 := "hahahaha" // string 底层就是一个byte的数组
     n6 := n5[:]
     fmt.Println("string",n6)
    
     //修改string类型的文字 注意中文是占三个字节的用   []rune 转
     n7 := []byte(n5)
     n7[0] = 'x'  // 单引号是byte类型  双引号是string类型
     n5 = string(n7)
     fmt.Println(n5) 
    

     

 

    

}