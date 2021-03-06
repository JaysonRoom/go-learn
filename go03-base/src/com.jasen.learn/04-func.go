package main

import "fmt"

func main() {

	// 1、函数 无参数 有参数 有返回值
	sayHi()
	name := "Tom"
	b, _ := sayHi2(name)
	fmt.Printf("%s say:%s\n", name, b)

	//2、变参
	println(sayHi3("sum: %d", 1, 2, 3))

	//3、匿名函数
	fun := func(m int, n int) {
		fmt.Printf("result:%d\n", m+n)
	}
	fun(4, 5)

	//4、延迟调用
	deferFunc(0)

	// 5、异常机制
    //--待补充--
}

//延迟调用
func deferFunc(x int) {
	//关键字 defer ⽤用于注册延迟调⽤用。这些调⽤用直到 ret 前才被执⾏行，通常⽤用于释放资源或错误处理
	//多个 defer 注册，按 FILO 次序执⾏行。哪怕函数或某个延迟调⽤用发⽣生错误，这些调⽤用依旧会被执⾏行
	defer println("a")
	defer println("b")
	defer func() {
		//recover()
		println(100 / x) // 除以0异常未被捕获，逐步往外传递，最终终⽌止进程。

	}()
	defer println("c")
}

// 无参数 无返回值
func sayHi() {
	println("Hi，Jasen")
}

// 有参数 有返回值
func sayHi2(x string) (b string, i int) {
	fmt.Printf("Hi，%s\n", x)
	return "bye", 1
}

//变参
func sayHi3(s string, n ...int) string {
	sum := 0
	for _, i := range n {
		sum += i
	}
	return fmt.Sprintf(s, sum)
}
