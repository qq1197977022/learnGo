package main

func main() {

}

/*
1.Unix基本哲学之一即一切皆文件, 可以执行 打开/ 关闭, 读/ 写操作
2.socket起源于Unix, 本质是一种文件描述符
3.socket数据传输是一种特殊的I/O
	1.获取socket文件描述符
	2.建立连接
	3.数据传输
	4.关闭连接
4.常用socket类型
	1.面向连接的数据流socket
		1.适用于面向连接的TCP应用
	2.无连接的数据包socket
		1.适用于无连接的UDP应用
*/
