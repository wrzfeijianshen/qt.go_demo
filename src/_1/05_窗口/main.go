package main

import (
	"os"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_MainWindow2()
	mw.MainWindow.Resize(200, 200)
	mw.MainWindow.Show()

	if true {
		f1 := func(painter *qtgui.QPainter) {
			painter.SetPen2(1)
			painter.DrawLine2(10, 10, 100, 100)
		}
		mw.Centralwidget.InheritInitPainter(f1)
	}

	app.Exec()
}
