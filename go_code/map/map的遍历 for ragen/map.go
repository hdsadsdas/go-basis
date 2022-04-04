package main
import "fmt"

func main(){

	//简单的遍历
     haha := map[string]string{
          "1" : "哈哈",
		  "2" : "嘻嘻",
		  "3" : "呵呵",
	 }  

	 for n1,n2 := range haha {  //n1 为名  ,n2为值

           fmt.Printf("n1=%v,n2=%v\n",n1,n2)
         
	 }
	 fmt.Println("haha的长度",len(haha))
         
	 fmt.Println("-----------------------------")

	 //较为复杂的遍历

	 xixi := make(map[string]map[string]string)
     
	 xixi["1"] = make(map[string]string)
	 xixi["1"]["name"] = "哈哈"
	 xixi["1"]["old"] = "18"

	 xixi["2"] = make(map[string]string)
	 xixi["2"]["name"] = "嘻嘻"
	 xixi["2"]["old"] = "19"
	 xixi["2"]["dizi"] = "十万八千里"
	 
  for n3,n4 := range xixi{ //n3 为名 因为值是map所以n4是map
    fmt.Println("学生的id",n3)
	fmt.Println("学生的信息",n4)
	for n5,n6 := range n4{ //n5 为n4的值  n6为n5的值 
   
		  fmt.Println("\tn5=",n5,"n6=",n6)

	}
  }

  fmt.Println("xixi的长度",len(xixi))
  fmt.Println("xixi里面的的长度",len(xixi["2"]))



}