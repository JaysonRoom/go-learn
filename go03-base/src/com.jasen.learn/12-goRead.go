package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"syscall"
)

func main() {
	//读取当前文件相对路径
	url, _ := os.Getwd()
	url = url +"/src/com.jasen.learn"
	fileInfoList, err := ioutil.ReadDir(url)
	if err != nil {
		log.Fatal(err)
	}
	if len(fileInfoList) <= 15 {
		for i := range fileInfoList {
			fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
			fileExt := path.Ext(fileInfoList[i].Name())
			if fileExt == ".doc" || fileExt == ".docx" {
				println("word file")
				fileInfo, _ := os.Stat(url+"/"+fileInfoList[i].Name())
				fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
				fileAttributes := fileSys.FileAttributes
				fmt.Println(fileAttributes)
			}
		}
	} else {
		println("超出文件数，请将文件设置少于10个")
	}
	//对word文件读取文件的信息
	fmt.Scanf("输入任意键退出....")
}
