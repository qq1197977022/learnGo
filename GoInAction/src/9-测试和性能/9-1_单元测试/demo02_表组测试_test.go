package main

import (
	"net/http"
	"testing"
)

const checkMark1 = "\u2713"
const ballotX1 = "\u2713"

func TestDownload1(t *testing.T) {
	var urls = []struct { //结构体切片
		url        string
		statusCode int //预期状态码
	}{ //初始化
		{
			"http://2cifang.com",
			http.StatusOK, //200
		},
		{
			"http://2cifang.com/abc",
			http.StatusNotFound, //404
		},
	}

	t.Log("Given the need to test downloading different content.")
	{ //显式块
		for _, u := range urls { //迭代
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{ //显式块
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.", ballotX1, err)
				}
				t.Log("\t\tShould be able to Get the url", checkMark1)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\"status. %v", u.statusCode, checkMark1)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v", u.statusCode, ballotX1, resp.StatusCode)
				}
			}
		}
	}
}

//相比基础单元测试优点: 不用为每个URL创建测试函数
