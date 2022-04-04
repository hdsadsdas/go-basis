package main

import (
	"fmt"
	"reflect" //反射的包
)

func re(b interface{}){

	//将要要反射的值传入
	//返回一个 reflect.Type类型
	rtyp := reflect.TypeOf(b)
    fmt.Println("rtyp= ", rtyp)	
 
	//获取到reflect.Value
	rval := reflect.ValueOf(b)

	n2 := 100 + rval.Int()// 通过reflect.Value里的Int()方法
                         // 返回int类型的值
	fmt.Println(n2)

	fmt.Printf("rval=%v   rval类型是%T\n",rval,rval)

	//将rval  转成interface{}
	iv := rval.Interface()
   //将interface{}通过断言转回int
   
   num := iv.(int)

   fmt.Println(num)

}

//结构体的反射演示

type jgt struct{
	name string
	age int
}

func jgtfs(b interface{}){
			//返回一个 reflect.Type类型
			rtyp := reflect.TypeOf(b)
			fmt.Println("rtyp= ", rtyp)	

			//获取到reflect.Value
			rval := reflect.ValueOf(b)

			kind := rval.Kind()//类别 :  结构体和其他被重命名的总称
			fmt.Println(kind)

			iv := rval.Interface() 
			fmt.Println("iv",iv)

			tes := iv.(jgt)
			fmt.Println(tes.name) 


}


func main (){


	// var n1 int = 10
	// re(n1)

	//定义一个结构体
	n2 := jgt{"哈哈",20}
    
	jgtfs(n2)

}
