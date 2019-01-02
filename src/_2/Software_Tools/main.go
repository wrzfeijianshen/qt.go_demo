package main

import (
	"os"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_MainWindow2()
	if true {
		// 加载qss
		qss := qtcore.NewQFile1("./main.qss")
		ret := qss.Open(qtcore.QIODevice__ReadOnly)
		if ret {
			str := qtcore.QString_FromUtf81(qss.ReadAll())
			mw.MainWindow.SetStyleSheet(str)
			qss.Close()
		}
	}
	mw.Slot()

	mw.MainWindow.Show()
	app.Exec()
}
