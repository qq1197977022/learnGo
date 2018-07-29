package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		/*标签将变为对应字段的属性, 空标签等价于无标签
		1.标签是提供字段元信息的一种机制, 将JSON文档和结构类型的字段一一映射起来
		2.若无标签, 编解码过程会试图以大小写无关的方式使用字段名进行匹配
		3.若无法匹配则对应结构类型的字段为其零值
		*/
		UnescapedURL      string `json:"unescapedUrl"`
		URL               string `json:"url"`
		VisibleURL        string `json:"visibleUrl"`
		CacheURL          string `json:"cacheUrl"`
		Title             string `json:"title"`
		TitleNoFormatting string `json:"titleNoFormatting"`
		Content           string `json:"content"`
	}

	// gResponse 包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 向 Google 发起搜索
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
