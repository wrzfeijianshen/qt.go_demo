package main

import (
	"fmt"
	_ "fmt"
	"os"

	"github.com/kitech/qt.go/qtrt"

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
	if true {
		// 工具栏
		qCheckbox := qtwidgets.NewQCheckBox(mw.Form)
		qCheckbox.SetText("三态复选框")
		qCheckbox.SetTristate(true)
		//qCheckbox.SetTristatep()

		qCheckbox.SetCheckState(qtcore.Qt__PartiallyChecked)
		qCheckbox.Move(50, 50)
		label := qtwidgets.NewQLabel(mw.Form, 0)
		label.Move(50, 80)
		label.SetFixedSize1(100, 30)
		label.SetText("Click CheckBox...")
		if qCheckbox.IsTristate() {
			fmt.Println("是三态状态")
		}
		qtrt.Connect(qCheckbox, "stateChanged(int)", func(iVal int) {
			if iVal == qtcore.Qt__Checked {
				label.SetText("选中了")
			} else if iVal == qtcore.Qt__PartiallyChecked {
				label.SetText("半选")
			} else {
				label.SetText("Unchecked")
			}
		})

	}

	mw.Form.Show()
	app.Exec()
}
