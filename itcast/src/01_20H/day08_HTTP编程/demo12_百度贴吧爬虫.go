package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg2 sync.WaitGroup

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	spiderEngine(1) //配置爬虫引擎
}

func spiderEngine(n int) {
	wg2.Add(n)

	for i := 0; i < n; i++ {
		go spider(i) //启动爬虫
	}

	wg2.Wait()
}

func spider(i int) {
	defer wg2.Done()

	fmt.Printf("爬取第 %d 页\n", i)
	pn := strconv.Itoa(i * 50)
	url := "https://tieba.baidu.com/f?kw=绝地求生&ie=utf-8&pn=" + pn
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	spiderData(res, pn) //数据处理
}
func spiderData(res *http.Response, pn string) {
	fmt.Printf("处理第 %s 页数据\n", pn)

	f, err := os.Create("D:\\workspace\\go\\learnGo\\itcast\\src\\01_20H\\day08_HTTP编程\\tempFile\\" + pn + ".html")
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(f, res.Body)
	f.Close()
	res.Body.Close()
}

/*
绝地求生吧
	https://tieba.baidu.com/f?kw=绝地求生&ie=utf-8&pn=0
	规律: 翻页+50
*/
