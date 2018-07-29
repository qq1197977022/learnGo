package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var wg3 sync.WaitGroup

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	start, end := user()
	spiderEngine1(start, end)
}
func user() (start, end int) {
retry:
	fmt.Print("起始页 start(>=1) = ")
	fmt.Scan(&start)
	fmt.Print("终止页 start(>1) = ")
	fmt.Scan(&end)
	if start >= 1 && end > start {
		return
	}
	fmt.Println("你是不是傻?")
	goto retry
}

func spiderEngine1(start, end int) {
	wg3.Add(end - start)

	for i := start; i < end; i++ {
		go spider1(i)
	}

	wg3.Wait()
}

func spider1(i int) {
	defer wg3.Done()

	fmt.Printf("爬取第 %d 页\n", i)
	pn := strconv.Itoa(i)
	url := "https://www.pengfu.com/xiaohua_" + pn + ".html"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	recByts := make([]byte, 1024*100)
	_, err = res.Body.Read(recByts)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	spiderRE(recByts, pn)
}
func spiderRE(recByts []byte, pn string) {
	fmt.Printf("处理第 %s 页数据\n", pn)

	expr := `<h1 class="dp-b"><a href="https://www.pengfu.com/content_\d+?_1.html" target="_blank">(?s:(.+?))</a>(?s:.+?)</h1>(?s:.+?)<div class="content-img clearfix pt10 relative">(?s:.+?)(?s:(.+?))</div>`
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatalln(err)
	}
	recStr := fmt.Sprintf("%s", recByts)
	recSli := re.FindAllStringSubmatch(recStr, -1)
	f, err := os.Create("D:\\workspace\\go\\learnGo\\itcast\\src\\01_20H\\day08_HTTP编程\\tempFile\\" + pn + ".txt")
	defer f.Close()
	for _, sli := range recSli {
		if err != nil {
			log.Fatalln(err)
		}
		f.WriteString(fmt.Sprintf("%s\n%s\n", sli[1], sli[2]))
		f.WriteString("=======================================================================================\n")
	}
}
