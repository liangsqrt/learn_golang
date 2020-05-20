package learn_file

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ReadPath() {
	fmt.Println("hello world!")
	path1, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//  C:\Users\liang\AppData\Local\Temp
	//  E:\goproject\test  编译后
	fmt.Println("path1  ", path1)

	path2, _ := os.Getwd()
	//  path2 E:\goproject\test
	//  path2 E:\   在e盘执行main.exe 得到的内容
	fmt.Println("path2  ", path2)

	path3, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	//  C:\Users\liang\AppData\Local\Temp\___1go_build_main_go.exe
	fmt.Println("path3  ", path3)

	path4, _ := os.Executable()
	path5 := filepath.Dir(path4)
	//  C:\Users\liang\AppData\Local\Temp
	fmt.Println("path5  ", path5)
}

func CreatePath() {
	path1, _ := os.Getwd()
	os.OpenFile(filepath.Join(path1, "learn_file"), os.O_RDONLY, 0666)
	//  什么鬼，2020年05月20日
}
