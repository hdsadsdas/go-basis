package main 
import "fmt"


type haha func(int, int) int//haha  就为函数类型

func xixi (n1 int,n2 int) (sum int) {

       sum = n1 + n2
   return

}

func hehe (leixing haha, n1 int, n2 int)(sum int , sub int , hanshu int){
     sum = n1 + n2
	 sub = n1 - n2
	 hanshu = leixing(n1,n2) 
	return  
}


func main (){
  
   type name int//相当于把int 取了个别名为name

   var n1 name = 10

   fmt.Println("n1为",n1)

   sum := xixi(10,20)

  fmt.Println("sum为",sum)

  sum2,sub2,hanshu := hehe(xixi,20,30)
  fmt.Println("sum2为",sum2)
  fmt.Println("sub2为",sub2)
  fmt.Println("hanshu为",hanshu)
}
