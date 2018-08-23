package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark2 = "\u2713" //unicode码 ~ 对号
	ballotX2   = "\u2717"
)

func init() {
	Routes1()
}

func TestSendJson(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil) //返回*Request
		if err != nil {
			t.Fatal("\tShould be able to create a request", ballotX2, err)
		}
		t.Log("\tShould be able to create a request", checkMark2)

		resRec := httptest.NewRecorder()
		/*
			1.返回一个初始化的ResponseRecorder
			2.ResponseRecorder是ResponseWriter的一个实现, 记录其冲突, 用于在后续测试中检查
		*/
		http.DefaultServeMux.ServeHTTP(resRec, req) //在默认ServMux中, 把HTTP请求分派给模式最接近请求URL的处理程序

		if resRec.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX2, resRec.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark2)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(resRec.Body).Decode(&u); err != nil {
			/*
				1.NewDecoder返回从resRec.Body中读取数据的*Decoder
					1.decoder自己实现buffer
					2.从resRec.Body中读取的数据可能超出解码JSON值所需数据
				2.Decode从其输入中读取下一个JSON-encode值并存入&u中
			*/
			t.Fatal("\tShould decode the response", ballotX2)
		}
		t.Log("\tShould decode the response", checkMark2)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name", checkMark2)
		} else {
			t.Error("\tShould have a Name", ballotX2, u.Name)
		}

		if u.Email == "bill@qq.com" {
			t.Log("\tShould have an Email", checkMark2)
		} else {
			t.Error("\tShould have an Email", ballotX2, u.Email)
		}
	}
}

func Routes1() {
	http.HandleFunc("/sendjson", SendJson1)
}

func SendJson1(rw http.ResponseWriter, r *http.Request) {
	//var name, email string
	//fmt.Println("请输入姓名和邮箱(以空格分隔):")
	//fmt.Scanln(&name, &email)

	u := struct {
		Name  string
		Email string
	}{
		"Bill",
		"bill@qq.com",
		//name,
		//email,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

/*
1.测试代码只能访问导出标识符
2.
*/
