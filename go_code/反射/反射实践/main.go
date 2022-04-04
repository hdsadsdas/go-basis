package main

import (
	"fmt"
	"reflect"
)


type Haha struct{
	n1 int
	n2 int
}

func (this Haha) Get(name string) string{

	fmt.Println(name,"操作",this.n1+this.n2)
	return name

}

//反射
func fs(n Haha){

	val := reflect.ValueOf(n)
	//typ := reflect.TypeOf(n)

    //通过反射查看有几个字段
	num := val.NumField()
    fmt.Println("haha有几个字段",num)

	for i := 0 ; i < num ; i++{

		fmt.Printf("haha的第%d个字段为%v\n",i,val.Field(i))

	} 

	//调用haha的方法
	
	var xixi []reflect.Value
	xixi = append(xixi,reflect.ValueOf("哈哈"))
	res := val.Method(0).Call(xixi)
	fmt.Println(res[0].String())
}


func main(){

   
   xixi := Haha{1,2}
  
   fs(xixi)
  
}