package main

import (
   "fmt"
   "encoding/json"
)

type haha struct {
	name string  
	id int
	age int
}

func text() {

	xixi := haha{
		name : "哈哈",
		id : 20000,
		age : 20 ,
	}

//将上面的结构体变量序列化

    jieshou , ok := json.Marshal(xixi) //返回一个byte切片

	if ok != nil {
		fmt.Println("序列化错误")
		return
	}
	
	fmt.Println("序列化后",string(jieshou))


}
//map的序列化
func textmap() {

	var haha map[string]interface{}

	haha = make(map[string]interface{})

	haha["名字"] = "嘻嘻"
	haha["age"] = 29
	haha["id"] = 123232 

	//将map序列化

	jieshou , ok := json.Marshal(haha) //返回一个byte切片  序列化后的数据

	if ok != nil {
		fmt.Println("序列化错误")
		return
	}
	
	fmt.Println("序列化后",string(jieshou))


}


//反序列化
func fan(){

	xixi := "{\"age\":29,\"id\":123232,\"名字\":\"嘻嘻\"}"

	var haha map[string]interface{}

	//反序列化
	//传入你要反序列化的byte切片  和序列化前的类型(地址)
    //反序列化map 不需要make 因为json.Unmarshal这个方法已经make好了

	ok := json.Unmarshal([]byte(xixi),&haha)
	if ok != nil {
           fmt.Println("反序列化失败")
		   return
	}

	fmt.Println("反序列化后",haha)


}



func main(){

   text()
   textmap()
   fan()
}