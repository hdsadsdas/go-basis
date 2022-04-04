package main
import "fmt"


func main(){

    var name string
	var age byte
	var hh float32
	var xx bool


	fmt.Println("输入姓名")
	fmt.Scanln(&name)

    fmt.Println("输入年龄")
	fmt.Scanln(&age)

	fmt.Println("输入工资")
	fmt.Scanln(&hh)

    fmt.Println("输入是否通过")
	fmt.Scanln(&xx)

    fmt.Printf("姓名是%v \n 年龄是%v \n 工作是%v \n 是否通过%v \n",name,age,hh,xx)
  
}