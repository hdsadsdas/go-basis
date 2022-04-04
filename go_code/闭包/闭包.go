package main
import "fmt"

func xixi() func(int) int{
    

	//与 n1 构成了闭包，下次调用会改变这个值
	var n1 int = 10
  
	 return func(n2 int) int {
		     n1 = n1 + n2 
			 return n1
	 }
}

func haha(ming string) func(string) string{

	//与ming 不构成闭包
     return func(mak string) string{
           
		  ming := ming + mak 
          return ming
	 }
  
}

func main(){
 
  n1 := xixi()

   fmt.Println("n1=",n1(1))
   fmt.Println("n1=",n1(2))
   fmt.Println("n1=",n1(3))

  //不构成闭包
  bi := haha("天")
  fmt.Println("bi="+bi("空"))
  fmt.Println("bi="+bi("气"))



}