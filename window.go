package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

type Window interface {
	ShowWindow()//展示窗口信息
}

//创建压缩解压缩的界面类
type ComWindow struct {
	Window
}

//创建提示信息的界面类
type LabelWindow struct {
	Window
}

//创建界面类对象
func Show(Window_Type string)  {
	var Win Window
	switch Window_Type {
	case "main_window" :
		Win = &ComWindow{}
	case "lab_window" :
		Win = &LabelWindow{}
	default:
		fmt.Println("参数传递错误")
	}
	Win.ShowWindow()
}

var Text string

//实现ShowWindow方法，展示出空白的窗口
func (Com *ComWindow)  ShowWindow() {

}

func (cw *ComWindow) StartToUnZip(file string, saveFile string) {
	reader, err := zip.OpenReader(file)
	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	for _, file := range reader.File {
		rc, err := file.Open()//打开的是文件或者文件夹
		if err != nil {
			fmt.Println(err)
		}
		defer rc.Close()

		newName := saveFile + file.Name

		//判断是否是文件夹
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(newName, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			//是文件
			f, err := os.Create(newName)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			_, err1 := io.Copy(f, rc)
			if err1 != nil {
				fmt.Println(err1)
			}
		}
	}
}

func (lab *LabelWindow) ShowWindow() {

}

func (mv *ComWindow) StartToZip(filePath string, savePath string)  {
	d, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()

	//打开该文件
	file, err := os.Open(savePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fileHeader, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println(err)
	}
	//将要压缩的文件写入压缩包中
	w := zip.NewWriter(d)
	defer w.Close()
	writer, err := w.CreateHeader(fileHeader)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(writer, file)
}
