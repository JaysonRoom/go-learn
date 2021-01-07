# 接口
## 1、内容
- [x] 定义
- [x] 接口与类型关系
- [x] 接口嵌套
- [x] 空接口

## 2、学习笔记

![img](/imgs/07-interface.png)

## 3、Demo练习

```
package main

import "fmt"

func main() {
	var x Sayer
	dog := animal{name: "xiaogou"}
	cat := animal{name: "xiaomi"}
	x = dog
	x.sayHi()
	x = &cat
	x.sayHi()

	var y Sayer
	//person1 := people{name: "Jasen"}
	person2 := people{name: "Tom"}
	//y = person1
	//y.sayHi()

	y = &person2
	y.sayHi()

	//案例2 
	var peo People = &Student{}  // var peo People = Student{}将会编译不通过
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// 案例2
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}



// 案例1
type Sayer interface {
	sayHi()
}

type animal struct {
	name string
}

func (dog animal) sayHi() {
	fmt.Printf("%s colud say wawawa\n", dog.name)
}

type people struct {
	name string
}

func (pp *people) sayHi() {
	fmt.Printf("%s colud say hi\n", pp.name)
}

```

##### 参考资料
- [go学习笔记](https://github.com/qyuhen/book)