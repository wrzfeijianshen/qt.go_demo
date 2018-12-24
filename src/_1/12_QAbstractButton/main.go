package main

import (
	"fmt"
	"os"

	_ "github.com/kitech/qt.go/qtgui"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()

	btn := qtwidgets.NewQAbstractButton(nil)
	if true {

	}
	btn.Move(100, 100)
	btn.Resize(300, 100)
	btn.SetText("aaaaa")
	btn.SetWindowTitle("aaaBB")
	btn.Show()
	btn.SetDisabled(false)

	if true {
		// 加载qss
		qss := qtcore.NewQFile1("./main.qss")
		ret := qss.Open(qtcore.QIODevice__ReadOnly)
		if ret {
			str := qtcore.QString_FromUtf81(qss.ReadAll())
			mw.Form.SetStyleSheet(str)
			qss.Close()
		}
	}
	qtrt.Connect(btn, "clicked(bool)", func(bool) {
		fmt.Println("点击了")
	})

	mw.Form.Show()
	app.Exec()
}
