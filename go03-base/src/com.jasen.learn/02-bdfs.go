package main

import "fmt"

func main() {
	/** 1、复合对象初始化 **/
	//初始化复合对象，必须使⽤用类型标签，且左⼤大括号必须在类型尾部。
	// var a struct{ x int } =	{ 100 }  // error
    var a  = struct{ x int }{ 100 }

    fmt.Printf("A:%d\n",a.x)

    /**2、控制流**/
    x := 0
    if n := "abc"; x>0 {
    	println(n[2])
	}else if x<0 {
		println(n[1])
	}else {
		println(n[0])
	}

	/**3、for三种循环模式**/
	s := "abc"
	l := len(s)
	for i := 0; i< l;i++{
		println(i)
	}

	for l>0 {              //代替 while(n>0)
		println(s[l-1])
		l--
	}
	//for {   //代替while(true)

	//}
    /**4、Range**/
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m { // 返回 (key, value)。  //其中可忽略不想要的返回值k,v，或⽤用 "_" 这个特殊变量。
		println(k, v)
	}
}
