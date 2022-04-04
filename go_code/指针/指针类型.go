package main
import "fmt"

func main()  {
	 
    var i int =100

    var j *int =&i
   
	fmt.Println("i的地址为",&i)
    fmt.Println("j指针指向的地址为",j)
    fmt.Println("指针所在的地址为",&j)

   *j = 100+100//可以通过 *j 来修改 i 的值
    fmt.Println("i为",i)
    fmt.Println("*j为",*j)
    
}