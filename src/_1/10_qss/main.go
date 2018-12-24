package main

import (
	_ "fmt"
	"os"

	"github.com/kitech/qt.go/qtcore"

	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
	mw.Form.Resize(500, 500)
	if false {
		// 加载qss
		qss := qtcore.NewQFile1("./main.qss")
		ret := qss.Open(qtcore.QIODevice__ReadOnly)
		if ret {
			// str := qtcore.NewQString()
			// a := qss.ReadAll()
			// str := str.Operator_equal5(a)
			// mw.Form.SetStyleSheet(str)
			str := qtcore.QString_FromUtf81(qss.ReadAll())
			mw.Form.SetStyleSheet(str)
			qss.Close()
		}
	}

	mw.Form.Show()
	app.Exec()
}
