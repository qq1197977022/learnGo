package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

//feed模拟期望接收的XML文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
 <rss>
 <channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>http://www.goinggo.net/</link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
		<title>Object Oriented Programming Mechanics</title>
		<description>Go is an object oriented language.</description>
		<link>http://www.goinggo.net/2015/03/object-oriented</link>
	</item>
 </channel>
 </rss>`

/*
 1.mockserver返回用于处理请求的服务器的指针
 2.httptest.Server是一个HTTP服务器
 	1.在本地循回接口监听系统选择的端口, 用于端对端HTTP测试
*/
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

/*
1.为解决测试过程依赖资源的不可控性, 对测试持续集成和部署的自动化的影响, 需要进行模拟测试
2.Go标准库对模拟模拟测试的支持
	1.httptest包对HTTP测试提供了实用工具
		eg. 1.模拟基于HTTP网络的调用(模拟互联网的请求和响应)
*/
