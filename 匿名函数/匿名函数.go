package main
import "fmt"

//全局匿名函数
    var (
		//fun1 就是全局匿名函数
		Fun1 = func(n1,n2 int) int {
			return n1 * n2
		}
	)


func main(){

    sum := func (n1,n2 int) int{

         return n1+n2
       

	 }(10,20)

 fmt.Println("sum=",sum)
    
    sub := func (n1,n2 int) int {

           return n1-n2

	}
	
	fmt.Println("sub=",sub(10,20))

    fmt.Println("cheng=",Fun1(10,10))



}