package day05

/*
参见
	1.https://golang.org/cmd/go/
	2.Ctrl + F 搜索tool
		1.Run specified go tool
Demo
	1.生成CPU profile: go test -bench=. -cpuprofile cpu.out
		1.查看画像
			1.交互模式
				1.go tool pprof fabonacci.test.exe cpu.out
					1.top10 查看top N
			2.web
				1.go tool pprof -http=:80 cpu.out
				2.go tool pprof -web cpu.out
		2.输出画像到文件
			1.go tool pprof -png cpu.out > cpu.png
			2.go tool pprof -pdf cpu.out > cpu.pdf
*/
//13:00
