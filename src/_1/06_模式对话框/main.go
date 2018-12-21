package main

import (
	"os"

	"github.com/kitech/qt.go/qtcore"

	_ "github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_MainWindow2()
	mw.MainWindow.Resize(500, 500)
	mw.MainWindow.Show()
	pDialog := qtwidgets.NewQDialog(nil, 0)
	pDialog.Resize(400, 400)
	pDialog.SetWindowTitle("子对话框")
	pDialog.SetModal(true) // 设置模态
	pDialog.Show()

	pWidget := qtwidgets.NewQWidget(nil, 0)
	pWidget.Resize(300, 300)
	pWidget.SetWindowTitle("子窗口")
	pWidget.SetWindowModality(qtcore.Qt__ApplicationModal) // 设置模态
	pWidget.Show()

	app.Exec()
}
