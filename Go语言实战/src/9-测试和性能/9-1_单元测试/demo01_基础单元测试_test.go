package main

// 这个示例程序展示如何写基础单元测试
import (
	"net/http"
	"testing"
)

const checkMark = "\u2713" //√
const ballotX = "\u2717"   //×

// TestDownload 确认 http 包的 Get 函数可以下载内容
func TestDownload(t *testing.T) {
	//url := "http://2cifang.com"
	url := "https://www.baidu.com/"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}

			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

/*
1.测试对象
	1.包
	2.程序的一部分
	3.函数
2.目的: 验证目标代码能否在给定场景达到预期效果
3.测试场景
	1.正向测试: 正常执行情况下, 保证代码不产生错误
		eg. 确认代码是否可以成功向数据库中插入一条记录
	2.负向测试: 保证代码不仅产生错误, 且是预期的错误
		eg. 可能原因: 对数据库查询时没有找到任何结果, 或对数据库做了无效更新
3.单元测试
	1.基础测试: 只使用一组参数和预期结果来测试一段代码
	2.表组测试: 		多
*/
/*
1.Go语言测试工具
	1.testing包对Go package的自动化测试提供了支持
		1.提供从测试框架到报告测试的输出和状态的各种测试功能的支持
	2.测试框架识别测试对象约定
		1.测试文件
			1.必须以 _test.go 结尾
		2.测试函数
			1.签名: 必须公开 ~ 导出
				1.格式: TestXxx ~ 1.以Test开头, 自定义部分Xxx必须首字母大写 ~ 大驼峰
			2.形参: 必须是test.T类型的指针
				1.T类型管理测试状态和格式化的测试日志
			3.返回值: 无
*/
