package main

func main() {

}

/*
一.runner包: 管理处理任务的运行和生命周期
	1.使用通道来监视程序执行时间, 也可用于终止程序
	2.当需要调度后台处理任务程序时, 该模式很有用
	3.该程序可能作为cron作业执行或在基于定时任务的云环境(eg.iron.io)里执行
*/
