package main

import (
	_ "fmt"
	"os"

	"github.com/kitech/qt.go/qtgui"

	_ "github.com/kitech/qt.go/qtgui"

	"github.com/kitech/qt.go/qtcore"
	_ "github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
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
	if false {
		// 工具栏
		qtoolbtn := qtwidgets.NewQToolButton(mw.Form)
		//qtoolbtn.SetArrowType(qtcore.Qt__DownArrow)
		qtoolbtn.SetArrowType(qtcore.Qt__LeftArrow)

		qtoolbtn.SetText("left arrow")
		qtoolbtn.SetToolButtonStyle(qtcore.Qt__ToolButtonTextUnderIcon)
		//qtoolbtn.SetStyleSheet("QToolButton{border: none; background: rgb(68, 69, 73); color: rgb(0, 160, 230);}")

	}

	if false {
		qAction := qtwidgets.NewQAction(mw.Form)

		qAction.SetText("工具按钮")
		icon := qtgui.NewQIcon2(":/images/logo1.png")
		qAction.SetIcon(icon)
		qtoolbtn := qtwidgets.NewQToolButton(mw.Form)
		qtoolbtn.SetIconSize(qtcore.NewQSize1(48, 48))
		qtoolbtn.Move(100, 100)
		qtoolbtn.SetFixedSize1(220, 100)

		qtoolbtn.SetDefaultAction(qAction)
		qtoolbtn.SetToolButtonStyle(qtcore.Qt__ToolButtonTextUnderIcon)
		qAction.SetToolTip("战斗不止")
	}

	mw.Form.Show()
	app.Exec()
}
