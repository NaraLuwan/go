package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 获取到当前目录，相当于python里的os.getcwd()
	pwd, _ := os.Getwd()
	fmt.Println("当前的操作路径为:", pwd)
	//文件路径拼接
	csvFile := filepath.Join(pwd, "code", "test.csv")
	fmt.Println("文件的路径为:", csvFile)
	//准备读取文件
	fs, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	// 一行一行读取
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}
}
