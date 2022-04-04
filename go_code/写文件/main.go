package main

import (
	"fmt"
    "os"
	"bufio"
	"io/ioutil"
)

func n1(){
		// func OpenFile(name string, flag int, perm FileMode) (*File, error) {    
	// 	...                   
	// }                       形参分别未  文件名  os里的不同操作   随便的4位数


	//返回一个文件句柄

	haha, ok := os.OpenFile("./haha.txt",os.O_CREATE | os.O_TRUNC | os.O_WRONLY ,4423)

	defer haha.Close()

	if ok != nil {
		fmt.Println("打开错误")
		return
	}
  

	//write 写入的方法 这个是写入字节
	//传入一个byte类型的切片
	haha.Write([]byte("zhe shi yi ge qie pian\n"))

	//writeString 写入的是字符串
	haha.WriteString("这是字符串\n")
}

func n2(){

	//返回一个文件句柄
 haha, ok := os.OpenFile("./haha.txt",os.O_CREATE | os.O_APPEND | os.O_WRONLY ,4423)

 defer haha.Close()

 if ok != nil {
	 fmt.Println("写入错误")
	 return
 }

 //创建一个写的对象
 wr := bufio.NewWriter(haha)
 wr.WriteString("这是第二种方法")//将数据先写入缓存

 wr.Flush()//将缓存写入文件中

}

func n3(){


	str := "hello 沙河"
	err := ioutil.WriteFile("./haha.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	
     }
}


func main(){


  n1()
	
  n2()

  n3()

}