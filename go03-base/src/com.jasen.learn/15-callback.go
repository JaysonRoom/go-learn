package main

import "fmt"

type callback func(x,y int) int
var defaultUsrAction  *userAction

type userAction struct {
	name string
	cbFunc []callback  //struct嵌入方法
}

func (usr *userAction)Init(){
	defaultUsrAction =&userAction{
		name:"funcCal",
		cbFunc: make([]callback,0),
	}
}

func (usr *userAction) register(str string,c callback)  {
	usr.name = str
	usr.cbFunc = append(usr.cbFunc,c)
}

func main() {
	x, y := 1, 2
	defaultUsrAction.Init() //初始化
	defaultUsrAction.register("add",add) //注册回调方法
	defaultUsrAction.register("mul",mul)
	//fmt.Println(testCb(x, y, add))

	delFunc(x,y ,defaultUsrAction)
	testCb(x,y,add) //调用方法
}

func delFunc(x,y int,dua *userAction)  {
	funds := dua.cbFunc
	for _, listens := range funds {
		listens(x,y)
	}
}

func testCb(x,y int,add callback) int {
	x = x+y
	return add(x,y)
}

func add(x,y int) int {
	fmt.Printf("add result is %d\n",x+y)
	return x+y
}
func mul(x,y int) int {
	fmt.Printf("mul result is %d\n",x-y)
	return x-y
}
