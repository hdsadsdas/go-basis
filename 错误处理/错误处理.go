package main
import "fmt"

func test(){

	//使用defer + recover 来捕获错误

	defer func(){
         
		 sum := recover() //recover()内置函数  捕获这里面的错误
         if sum != nil{ // nil判断错误
			 fmt.Println("sum=错误",sum)
		 }   

	}()


	n1 := 1
	n2 := 0

	n3 := n1/n2 //这是一个错误
	fmt.Println("n3=",n3)

}




func main(){

	test()

	fmt.Println("哈哈哈")


}