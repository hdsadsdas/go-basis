package main
import "fmt"

func main(){
   //第一种使用方式
	var a map[string]string
	//在使用map前，要用make 给 map 分配空间 不然map没有空间
	a = make(map[string]string,10)
	a["1"] = "哈哈"
	a["2"] = "嘻嘻"
	a["2"] = "叽叽" // key 是关键字 重复会覆盖
  
	fmt.Println("第一种方式",a)

	//第二种
   xixi := make(map[int]string,10)
   xixi[1] = "哈哈"
   xixi[2] = "嘻嘻"
   fmt.Println("第二种方式",xixi)
   
   //第三种
   
   haha := map[string]int{
        "哈哈" : 1,
		"嘻嘻" : 2,

   }
   fmt.Println("第三种方式",haha)

   //使用
   lala := make(map[string]map[string]string)

   lala["01"] = make(map[string]string)//因为值是map使用要make
   lala["01"]["name"] = "哈哈"
   lala["01"]["old"] = "18"
   
   lala["02"] = make(map[string]string)//因为值是map使用要make
   lala["02"]["name"] = "嘻嘻"
   lala["02"]["old"] = "19"

   fmt.Println("使用",lala)
   fmt.Println("使用2",lala["01"])
   fmt.Println("使用3",lala["01"]["name"])

   //删除
   delete(lala["01"],"name")
   fmt.Println("删除",lala)
  
   //清空
    lala = make(map[string]map[string]string) //  指向了一个新的空间 类型要和前面相同

   fmt.Println("清空",lala)

   //map的查找

   n1,n2 := a["1"]  // n1 为找到的值 n2为返回值:对或错

   if n2{
	   fmt.Println("找到了值为",n1)
   }else{
	   fmt.Println("没找到")
   }


}
