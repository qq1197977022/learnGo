package main

/*
1.日志是查找潜在bug, 了解程序工作状态的方式
	1.可用于跟踪 & 调试 & 分析代码
2.Go标准库提供了log包, 可对日志做最基本的配置, 开发人员还可自定义日志记录器
3.传统CLI(命令行界面)程序直接将输出写到名为stdout设备上, 所有操作系统都有这种设备, 该设备默认目的地是标准文本输出
	1.默认设置下, 终端会显示写到stdou设备上的文本, 这对于单个目的的输出很方便
	2.有事会需要同时输出程序信息和执行细节
		1.执行细节称为日志
	3.为解决程序输出和日志混为一谈, UNIX架构上增加了名为stderr的设备, 该设备为日志默认目的地
	4.若想在程序运行期间同时看到程序输出和日志, 可将终端配置为同时显示写到stdout和stderr的信息
	5.如果程序没有输出, 通常是将日志信息写到stdout设备, 将错误或警告信息写到stderr设备
*/
