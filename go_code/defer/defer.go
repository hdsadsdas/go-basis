package main
import "fmt"

func haha(n1,n2 int) int {
	//defer 要等函数执行完毕后再执行，在ruturn前执行
    //按照先入后出的顺序执行
	defer fmt.Println("n1=",n1)//第三执行 拷贝入栈不会受到后面代码的影响
	 defer fmt.Println("n2=",n2)//第二执行  拷贝入栈不会受到后面代码的影响
    
	 n1++
	 n2++

	 sum := n1 + n2
	 fmt.Println("sum=",sum) // 最先执行
	 return sum
}



func main(){

    sum := haha(10,20)

   fmt.Println("mian sum = ",sum)//最后执行

}

