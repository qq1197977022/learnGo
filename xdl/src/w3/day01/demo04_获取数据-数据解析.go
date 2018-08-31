package day01

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DataFetchParse() {
	//url := "http://route.showapi.com/1196-2?showapi_appid=72961&showapi_sign=764746e5747045d28bd9e0855212e9e0&keyword=推波助澜"
	url := "http://route.showapi.com/1196-1?showapi_appid=72961&showapi_sign=764746e5747045d28bd9e0855212e9e0&keyword=男"
	if resp, err := http.Get(url); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatalln(err)
		}
	}
}
