package main

import (
	_ "fmt"
	"os"

	_ "github.com/kitech/qt.go/qtgui"

	"github.com/kitech/qt.go/qtcore"
	_ "github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()

	// 不成功....
	btn := qtwidgets.NewQPushButton(mw.Form)
	btn.SetText("主菜单")
	btn.Move(50, 50)

	qmenu := qtwidgets.NewQMenu(nil)
	qmenu.AddSection("设置")
	qmenu.AddSection("版本检测")
	qmenu.AddSeparator()
	qmenu.AddSection("关于我们")
	qmenu.AddSection("退出")
	qmenu.Move(100, 100)

	btn.SetMenu(qmenu)

	if true {
		// 加载qss
		qss := qtcore.NewQFile1("./main.qss")
		ret := qss.Open(qtcore.QIODevice__ReadOnly)
		if ret {
			str := qtcore.QString_FromUtf81(qss.ReadAll())
			qss.Close()

			mw.Form.SetStyleSheet(str)

		}
	}

	mw.Form.Show()
	app.Exec()
}
