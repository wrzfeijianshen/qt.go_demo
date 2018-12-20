package main

import (
	"os"

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
			//s := qcolor.NewForInherit()// 默认颜色
			s := qcolor.NewForInherit1(qcolor.Blue())
			painter.SetPen(s)
			painter.DrawLine2(10, 10, 100, 100)
		}

		mw.Form.InheritInitPainter(f1)
		//mw.Form.InheritInitPainter(f1)// 不能同时执行多次
	}

	eg = false
	if eg {
		// 多线颜色
		f1 := func(painter *qtgui.QPainter) {
			var qcolor qtgui.QColor
			for i := 0; i <= 30; i++ {
				s := qcolor.NewForInherit1(i)
				painter.SetPen(s)
				painter.DrawLine2(10, 10+i*10, 100, 10+i*10)
			}
		}
		mw.Form.InheritInitPainter(f1)
	}

	eg = false
	if eg {
		// 多线颜色 NewForInherit2 rgba
		f1 := func(painter *qtgui.QPainter) {
			var qcolor qtgui.QColor

			// [-2] void QColor(int, int, int, int)
			// 查帮助文档 QColor::QColor(int r, int g, int b, int a = ...)
			// QColor::QColor ( int r, int g, int b, int a = 255 )
			// func (*QColor) NewForInherit2(r int, g int, b int, a int) *QColor
			s := qcolor.NewForInherit2(255, 255, 255, 255)
			painter.SetPen(s)
			painter.DrawLine2(10, 10, 100, 10)
			s = qcolor.NewForInherit2(255, 0, 0, 255)
			painter.SetPen(s)
			painter.DrawLine2(10, 40, 100, 40)
			s = qcolor.NewForInherit2(0, 255, 0, 255)
			painter.SetPen(s)
			painter.DrawLine2(10, 70, 100, 70)
			s = qcolor.NewForInherit2(0, 0, 255, 255)
			painter.SetPen(s)
			painter.DrawLine2(10, 100, 100, 100)
		}
		mw.Form.InheritInitPainter(f1)
	}

	eg = false
	if eg {
		// 多线颜色 NewForInherit2p rgb
		f1 := func(painter *qtgui.QPainter) {
			var qcolor qtgui.QColor
			s := qcolor.NewForInherit2p(255, 255, 255)
			painter.SetPen(s)
			painter.DrawLine2(10, 10, 100, 10)
			s = qcolor.NewForInherit2p(255, 0, 0)
			painter.SetPen(s)
			painter.DrawLine2(10, 40, 100, 40)
			s = qcolor.NewForInherit2p(0, 255, 0)
			painter.SetPen(s)
			painter.DrawLine2(10, 70, 100, 70)
			s = qcolor.NewForInherit2p(0, 0, 255)
			painter.SetPen(s)
			painter.DrawLine2(10, 100, 100, 100)
		}
		mw.Form.InheritInitPainter(f1)
	}
	app.Exec()
}
