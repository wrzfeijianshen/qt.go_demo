package main

import (
	"fmt"
	"os"

	"github.com/kitech/qt.go/qtcore"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
	mw.Form.Show()

	// QResizeEvent 事件
	if false {
		f1 := func(event *qtgui.QResizeEvent) {
			x := mw.Form.Width()
			y := mw.Form.Height()
			fmt.Printf("窗口发生改变 : x :[%d],y:[%d]\n", x, y)
		}
		// QResizeEvent
		// Inherit + QResizeEvent
		mw.Form.InheritResizeEvent(f1)
	}

	// mousePressEvent 鼠标事件
	if false {
		f1 := func(event *qtgui.QMouseEvent) {
			if event.Button() == qtcore.Qt__LeftButton {
				fmt.Println("鼠标左键按下")
			} else if event.Button() == qtcore.Qt__RightButton {
				fmt.Println("鼠标右键按下")
			} else {
				fmt.Println("鼠标其他键按下")
			}
		}
		mw.Form.InheritMousePressEvent(f1)
	}

	// QEvent 键盘事件
	if false {
		f1 := func(event *qtcore.QEvent) bool {
			if event.Type() == qtcore.QEvent__KeyPress {
				fmt.Println("键按下了")

				// key1 := qtgui.QInputEvent{event}
				// key := qtgui.QKeyEvent{&key1}

				// 继承QEvent <- QInputEvent <- QKeyEvent
				key := qtgui.QKeyEvent{&qtgui.QInputEvent{event}}
				if key.Key() == qtcore.Qt__Key_Tab {
					fmt.Println("tab键按下")
				}
			}
			return true
		}

		mw.Form.InheritEvent(f1)
	}

	// QMouseEvent 鼠标移动事件
	if true {
		f1 := func(event *qtgui.QMouseEvent) {
			//fmt.Println("鼠标来了111111: ", event.Type())
			if event.Type() == qtcore.QEvent__MouseMove {
				// 操作,鼠标左键按住窗体里面,左键移动不松开
				fmt.Println("鼠标在窗体里面移动")
			}
		}
		mw.Form.InheritMouseMoveEvent(f1)
	}

	app.Exec()
}
