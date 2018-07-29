package 待分类

import (
	"fmt"
	"os"
)

var path string

func main() {
	f := DirInput()

	if f != nil {
		fInfo, err := f.Stat() //获取FileInfo

		if err != nil {
			fmt.Printf("main-001\t%v\n", err)
		} else if fInfo.IsDir() {
			fun1(f)
		}

		f.Close()
	}
}

func DirInput() *os.File {
	//用户交互,获取目录路径
	fmt.Print("请输入要遍历目录路径: ")
	fmt.Scan(&path)

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("DirInput-001\t%v\n", err)
	}
	return f
}

//递归遍历
func fun1(f *os.File) {
	if f == nil {
		fmt.Println("参数f不能为nil")
		return
	}

	fInfoSlice, err := f.Readdir(0) //获取目录相关资源
	if err != nil {
		fmt.Printf("fun1-001\t%v\n", err)
	}

	if len(fInfoSlice) == 0 {
		fmt.Printf("%s是空目录\n", f.Name())
	} else {
		var nameSlice []string             //存储文件夹相关资源名称
		var dirPathSlice []string          //存储当前层子文件夹路径
		for _, fInfo := range fInfoSlice { //遍历当前目录,找出所有文件名和当前层子目录
			nameSlice = append(nameSlice, fInfo.Name())
			if fInfo.IsDir() {
				dirPathSlice = append(dirPathSlice, path+"\\"+fInfo.Name())
			}
		}
		fmt.Printf("%v--->目录内容\n\t%v\n", f.Name(), nameSlice)

		for _, dirPath := range dirPathSlice { //遍历当前层子目录
			f, err = os.Open(dirPath)
			if err != nil {
				fmt.Printf("fun1-002\t%v\n", err)
			}
			path = dirPath //遍历当前层内所有字目录前,修改path为父层路径
			fun1(f)
			f.Close()
		}
	}
}
