package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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
}
