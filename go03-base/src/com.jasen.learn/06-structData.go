package main

import "fmt"

type CallMethod interface {
	addAge()
	sayJasen(x string)
}

type UserInfo struct{
	name string
	age int
	sex int
	cmehod CallMethod
}
// 使用指针方式定义结构体方法
func (usr *UserInfo) addAge()  {
	usr.age +=1
}

func main()  {
	usr := UserInfo{}
	usr.name ="Jhone"
	usr.age = 12

	usr.cmehod.addAge()
	//usr.cmehod.sayJasen("Anti")
	fmt.Printf("age is %d\n",usr.age)
	usr.addAge()
	fmt.Printf("age is %d\n",usr.age)
}
