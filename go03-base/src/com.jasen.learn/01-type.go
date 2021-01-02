package main

import (
	"fmt"
)


func main() {
	/**1、定义变量**/
	var x string       // 自动初始化为零值或空值
	var y string = "A" // 常规初始化
	var z = "B"        // 省略变量类型
	t := "C"           // 局部定义变量
	var m, n = 1, "D"  // 一次定义多变量
	//多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 // (i = 0) -> (i = 2), (data[0] = 100)

	fmt.Printf("x:%s,y:%s, z:%s, t:%s, m:%d, n:%s\n", x, y, z, t, m, n)
	//输出结果：x:,y:A, z:B, t:C, m:1, n:D
	_, k := test()
	fmt.Printf("k:%s\n", k)
	// 输出结果 k:E


	/**2、常量**/
	// 关键字 iota 定义常量组中从 0 开始按⾏行计数的⾃自增枚举值。
	const sex = "男"
	const ltype,stype = "星期","日期"
	/**3、枚举**/
	const (
		sun = iota
		mon
		tues
		third
	)
	fmt.Printf("mon:%d\n",mon)
	//输出结果 mon:1

	/**4、字符串**/
    //4.1、使用索引号访问字符
    s1 := "abc"
    fmt.Println(s1[0] == 'a',s1[0] == 'b',s1[0] == 0x63)
    //输出结果： true true true
    //4.2、使用“`”定义不做转移处理的原始字符串，支持跨行
    s2 := `a
     b\n\rx00
     c`
    fmt.Printf(s2)
    //4.3、连接跨行字符串时，必须在上一行末尾用“+” 连接
    s3 := "hello "+
    	 "Jasen!"
	fmt.Printf(s3)
    //4.4、支持使用2个索引号返回子串，⼦子串依然指向原字节数组，仅修改了指针和⻓长度属性
    s4 :="hello Jasen!"

    s5 := s4[:5]
    s6 := s4[7:]
    s7 := s4[1:5]
    fmt.Printf("\ns5:%s,s6:%s,s7:%s\n",s5,s6,s7)
     //输出结果：s5:hello,s6:asen!,s7:ello
     //4.4、单引号字符常量表⽰示 Unicode Code Point，⽀支持 \uFFFF、\U7FFFFFFF、\xFF 格式。
	//对应 rune 类型，UCS-4。
	fmt.Printf("%T\n", 'a')  //单引号 类型是 int32
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")
    //输出结果:
    //int32
	//true true
	//4.5、修改字符串可以将其转换成 []rune 或 []byte，再转回string
	//4.6、for循环遍历字符串，有rune和byte2种方式
	s8 := "abcd汉字"
	for i :=0;i<len(s8);i++ {
		fmt.Printf("%c,", s8[i])
	}
	fmt.Println()
	//
	for _, r := range s8 { // rune
		fmt.Printf("%c,", r)
	}
	fmt.Println()
	//输出结果
	//a,b,c,d,æ,±,,å,­,,
	//a,b,c,d,汉,字,
	/**5、指针**/

}

func test() (i int, int string) {
	return 1, "E"
}
