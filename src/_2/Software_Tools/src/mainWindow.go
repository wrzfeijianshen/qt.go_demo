package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	. "internal/common/iobase"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func (this *Ui_MainWindow) Slot() {
	// 打开文件夹
	qtrt.Connect(this.BtnOpen, "clicked(bool)", func(bool) {
		//dir := qtcore.NewQDirp()// 打开当前文件夹
		// path := dir.CurrentPath()
		path := this.LineEdit.Text()

		path = strings.Replace(path, "/", "\\", -1)
		qpress := qtcore.NewQProcessp()
		qpress.StartDetached3("explorer " + path)
	})

	// 读取hosts文件
	qtrt.Connect(this.BtnReadHosts, "clicked(bool)", func(bool) {
		// 清空
		this.PlainTextEdit.Clear()

		path := this.LineEdit.Text()
		path = strings.Replace(path, "/", "\\", -1)
		path += "\\hosts"

		// qt
		if false {
			filename := qtcore.NewQFile1(path)
			if filename.Open(qtcore.QIODevice__ReadOnly | qtcore.QIODevice__Text) {
				fmt.Println("open ok")
				textStream := qtcore.NewQTextStream1(filename)
				ret := textStream.AtEnd()
				fmt.Println(ret) // 这里程序有bug?

				for !textStream.AtEnd() {
					fmt.Println("textStream AtEnd")
					str := textStream.ReadLinep()
					fmt.Println(str)
					this.PlainTextEdit.AppendPlainText(str)
				}
			} else {
				fmt.Println("open err")
				qtwidgets.NewQMessageBox1p(1, "err", "open file")
			}
			defer filename.Close()
		}

		// go
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(b)
		str := string(b)
		this.PlainTextEdit.AppendPlainText(str)
		// fmt.Println(str)

	})

	// 写hosts文件
	qtrt.Connect(this.BtnSaveHosts, "clicked(bool)", func(bool) {
		// 清空
		path := this.LineEdit.Text()
		path = strings.Replace(path, "/", "\\", -1)
		path += "\\hosts"

		// go
		str := this.PlainTextEdit.ToPlainText()
		fmt.Printf("str : [%s]", str)
		err := ioutil.WriteFile(path, []byte(str), 0777)
		if err != nil {
			fmt.Println("save file err")
			return
		}
	})

	// 编码转换 : 浏览 按钮
	qtrt.Connect(this.PushButton, "clicked(bool)", func(bool) {
		// https://blog.csdn.net/iamplane/article/details/69693807
		//dir := qtcore.NewQDirp()// 打开当前文件夹
		// path := dir.CurrentPath()
		path := this.LineEdit_2.Text()

		path = strings.Replace(path, "/", "\\", -1)
		fileFialog := qtwidgets.NewQFileDialog1(this.MainWindow, "打开文件", path, "*.*")
		str := fileFialog.GetExistingDirectoryp()
		// fmt.Println(str)
		this.LineEdit_2.SetText(str)
	})

	// 编码转换 : 查找 按钮
	qtrt.Connect(this.PushButton_3, "clicked(bool)", func(bool) {
		path := this.LineEdit_2.Text()
		path = strings.Replace(path, "/", "\\", -1)
		suffix := this.LineEdit_3.Text()
		if _, flag, err := IsExists(path); flag == 0 {
			// fmt.Println("输入文件夹存在")
		} else if flag == 1 {
			// fmt.Println("输入文件存在")
		} else if flag == -1 {
			fmt.Println("输入文件夹或文件不存在", err)
			return
		}
		var files []string
		var err error
		boxState := this.CheckBox.CheckState()
		if boxState == qtcore.Qt__Checked {
			// 如果选中
			files, err = WalkDir(path, suffix)
			if err != nil {
				return
			}
		} else {
			files, err = ListDir(path, suffix)
			if err != nil {
				return
			}
		}
		str := ""
		for _, fi := range files {
			str += fi
			str += "\n"

		}

		this.PlainTextEdit_2.SetPlainText(str)
	})

	// 编码转换 : 读取编码
	qtrt.Connect(this.PushButton_3, "clicked(bool)", func(bool) {

	})
}
