package main

import (
	"os"

	"github.com/kitech/qt.go/qtcore"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()

	eg := false

	// 默认代码
	if eg {
		f1 := func(painter *qtgui.QPainter) {
			painter.SetPen2(1)
			painter.DrawLine2(10, 10, 100, 100)
		}
		mw.Form.InheritInitPainter(f1)
	}
	mw.Form.Show()

	eg = true
	if eg {
		f1 := func(painter *qtgui.QPainter) {
			var qcolor qtgui.QColor

			s := qcolor.NewForInherit1(qtcore.Qt__red)

			painter.SetPen(s)
			painter.DrawLine2(10, 10, 100, 100)
		}
		mw.Form.InheritInitPainter(f1)
	}

	app.Exec()
}
