package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	Routes()

	log.Println("listener : Started : Listening on : 4000")
	http.ListenAndServe(":4000", nil)
	/*
		1.监听TCP网络地址, 然后通过handler调用服务来处理连接请求
		2.已接受连接被配置为启用TCP保持连接
		3.handler通常为nil; 此时, 使用默认ServeMux
	*/
}

//为网络服务设置路由 ~ 配置路由
func Routes() {
	http.HandleFunc("/sendjson", SendJson) //在默认ServeMux(HTTP请求多路复用器)中为给定模式注册处理函数	~ 配置路由: 将URL映射到对应的处理代码
}

//返回一个简单JSON文档
func SendJson(rw http.ResponseWriter, r *http.Request) {
	//用户交互
	var name, email string
	fmt.Println("请输入姓名和邮箱(以空格分隔):")
	fmt.Scanln(&name, &email)

	u := struct {
		Name  string
		Email string
	}{
		//"Bill",
		//"bill@qq.com",
		name,
		email,
	}

	rw.Header().Set("Content-Type", "application/json") //设置HTTP响应头键值对
	rw.WriteHeader(200)                                 //设置状态码
	json.NewEncoder(rw).Encode(&u)                      //1.返回向rw写入数据的*Encoder, 2.把&u中的JSON编码写入*Encoder
}

/*
1.服务端点(endpoint): 与服务宿主机信息无关, 用于分辨某服务地址, 一般不包含宿主的路径
2.构造网络API时, 通常希望直接测试自己的服务的所有服务端点, 而不用启动整个网络服务
*/
