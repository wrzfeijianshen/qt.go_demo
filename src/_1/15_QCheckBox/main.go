package main

import (
	"fmt"
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
	if false {
		// 工具栏
		qCheckbox := qtwidgets.NewQCheckBox(mw.Form)
		qCheckbox.SetText("三态复选框")
		qCheckbox.SetTristate(true) // 开启三态模型

		qCheckbox.SetCheckState(qtcore.Qt__PartiallyChecked)
		qCheckbox.Move(50, 50)
		label := qtwidgets.NewQLabel(mw.Form, 0)
		label.Move(50, 80)
		label.SetFixedSize1(100, 30)
		label.SetText("Click CheckBox...")
		if qCheckbox.IsTristate() {
			fmt.Println("是三态状态")
		}

		// 信号槽相应
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

	if true {

		// 设置布局
		m_pHBoxLayout := qtwidgets.NewQHBoxLayout1(mw.Form)
		m_pButtonGroup := qtwidgets.NewQButtonGroup(mw.Form)

		m_pButtonGroup.SetExclusive(true)
		for i := 0; i < 5; i++ {
			m_pRadioBtn := qtwidgets.NewQRadioButton(mw.Form)
			str := fmt.Sprintf("%s %d", "切换", i)

			m_pRadioBtn.SetText(str)
			m_pHBoxLayout.AddWidgetp(m_pRadioBtn)
			// m_pButtonGroup.AddButtonp(m_pRadioBtn)// 默认id从-2 -3开始
			m_pButtonGroup.AddButton(m_pRadioBtn, i)
		}
		m_pHBoxLayout.SetSpacing(10)
		m_pHBoxLayout.SetContentsMargins(10, 10, 10, 10)
		mw.Form.SetLayout(m_pHBoxLayout)
		// 信号槽相应
		qtrt.Connect(m_pButtonGroup, "buttonClicked(int)", func(iVal int) {
			str := fmt.Sprintf("%s %d", "切换", iVal)
			fmt.Println(str)
		})
	}

	mw.Form.Show()
	app.Exec()
}
