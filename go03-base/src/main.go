package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)


func main() {
	// 搭建服务
	fmt.Println("server startup http://localhost:9001  press ctrl+c stop!")
	
	http.HandleFunc("/", indexRouter) //设置访问的路由
	err := http.ListenAndServe(":9001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}	
}

func indexRouter(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Jasen!") //这个写入到w的是输出到客户端的
}