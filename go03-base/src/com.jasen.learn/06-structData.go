package main

import "fmt"

type UserInfo struct{
	name string
	age int
	sex int
}
// 使用指针方式定义结构体方法
func (usr *UserInfo) addAge()  {
	usr.age +=1
}

func main()  {
	usr := UserInfo{"Jasen",18,1}
	fmt.Printf("age is %d\n",usr.age)
	usr.addAge()
	fmt.Printf("age is %d\n",usr.age)
}
