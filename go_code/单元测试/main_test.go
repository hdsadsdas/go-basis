package haha

//要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，
//该文件包含 `TestXxx` 函数
import (
	"testing"
)

func TestAPP(t *testing.T){  //测试 一定要用 TestXxxx  格式 
	                        //用go test -v 启动
    
	res := App(10)

	if res != 55{
		t.Fatal("执行错误")   //testing的一个方法将输出错误和结束放一起
	}

	t.Logf("执行成功")//testing的一个方法将输出和结束放一起


}

func TestGetSub(t *testing.T){
	res := GetSub(10,3)

	if res != 7{
		t.Fatal("执行错误")   //testing的一个方法将输出错误和结束放一起
	}

	t.Logf("执行成功")//testing的一个方法将输出和结束放一起

}