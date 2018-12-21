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

	if false {
		f1 := func(painter *qtgui.QPainter) {
			var qcolor qtgui.QColor

			s := qcolor.NewForInherit1(qtcore.Qt__red)

			painter.SetPen(s)
			painter.DrawLine2(10, 10, 100, 100)
		}
		mw.Form.InheritInitPainter(f1)
	}

	if true {
		f1 := func(painter *qtgui.QPainter) {
			//func (this *QPainter) SetPen2(style int) style :0-6
			for i := 0; i < 6; i++ {
				painter.SetPen2(i)
				painter.DrawLine2(10, 10+i*20, 100, 10+i*20)
			}

			painter.SetPen2(qtcore.Qt__DashLine)
			painter.DrawLine2(10, 10+150, 200, 10+150)
		}
		mw.Form.InheritInitPainter(f1)
	}

	app.Exec()
}
