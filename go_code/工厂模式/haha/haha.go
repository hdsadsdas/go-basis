package haha


type Xixi struct{   // 首字母大写可以在别的包访问

   Name string
   Age int
}

type lala struct{  // 首字母小写不能在别的包访问
                   // 这时就要用到工厂模式
   name string
   age int

}

func Gcms(n string,a int) *lala{  //返回一个指针类型的 lala

   return &lala{  
	   name : n , //字段名首字母小写
	   age : a ,  //在别的包不能直接访问
	              // 可以用一个方法返回
   }

}

func (n *lala) Ff() string{

	return n.name

}

func (n *lala) Xg(nn string) {

    n.name = nn

}



