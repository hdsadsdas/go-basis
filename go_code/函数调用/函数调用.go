package main 
import "fmt"

func getSum(n1 int,n2 int) int {
   

	
    sum := n1 + n2
   fmt.Println("sum=",sum)
    return sum
}

func han(n1 int ){//没有返回值 在main中 只能调用函数 不能用变量来接收 

   n1 = n1 + 1 

   fmt.Println("han n1=",n1)
}


func he(n1 int, n2 int)(int,int){

sum := n1+n2
sub := n1-n2
return sum,sub


}


func main(){

   sum := getSum(10,20)
   fmt.Println("main sum = " ,sum)

   n1 := 1
    han(1)
   fmt.Println("main n1=", n1)
   
   he,cha := he(1,2) //多个返回值要用多个变量接收
   fmt.Println("和为=",he,"差为=",cha)

  // _,dan := he(3,2)//只像要一个返回值可以用下滑线占位
}


