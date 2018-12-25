package main

import (
	"os"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
	mw.Form.Resize(350, 350)
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

	qmenu := qtwidgets.NewQMenu(mw.Form)
	qmenu.AddAction("设置")
	qmenu.AddAction("版本检测")
	qmenu.AddSeparator()
	qmenu.AddAction("关于我们")
	qmenu.AddAction("退出")

	btn := qtwidgets.NewQPushButton(mw.Form)
	btn.SetText("主菜单")
	btn.SetFixedSize1(100, 30)
	btn.Move(50, 50)
	btn.SetMenu(qmenu)

	mw.Form.Show()
	app.Exec()
}
