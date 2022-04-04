package main
import "fmt"

//init 函数   先执行init 再执行main
func init(){

   fmt.Println("init")

}

func main(){

    fmt.Println("main")

}