package main

import (
	_ "fmt"
	"os"

	_ "github.com/kitech/qt.go/qtgui"

	"github.com/kitech/qt.go/qtcore"

	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
	if true {
		// 设置能够显示的位数
		mw.LcdNumber.SetDigitCount(25)

		// 设置显示的模式为十进制
		mw.LcdNumber.SetMode(qtwidgets.QLCDNumber__Dec)
		//mw.LcdNumber.SetMode(qtwidgets.QLCDNumber__Hex)

		// 设置显示外观
		mw.LcdNumber.SetSegmentStyle(qtwidgets.QLCDNumber__Flat)

		// 显示样式
		mw.LcdNumber.SetStyleSheet("border: 1px solid green; color: green; background: silver;")
		mw.LcdNumber.Display1(20191222)

		// 构建定时器
		// var timer *qtcore.QTimer
		// timer = timer.NewForInheritp()
		timer := qtcore.NewQTimer(nil)

		timer.SetInterval(1000)

		// qtrt.Connect(timer, "timeout()", func() {
		// 	fmt.Println("hehehhe222")
		// })
		qtrt.Connect(timer, "timeout()", func() {
			// 获取系统当前时间
			// var datime *qtcore.QDateTime
			// datime = datime.CurrentDateTimeUtc()
			datime := qtcore.QDateTime_CurrentDateTime()
			mw.LcdNumber.Display(datime.ToString1("yyyy-MM-dd HH:mm:ss.zzz"))
		})
		timer.Start1()

	}

	mw.Form.Show()
	app.Exec()
}
