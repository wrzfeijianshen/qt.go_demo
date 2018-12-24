package main

import (
	// "fmt"
	"os"

	_ "github.com/kitech/qt.go/qtcore"

	"github.com/kitech/qt.go/qtgui"

	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()

	if false {
		mw.Label.SetGeometry(10, 10, 200, 200)
		// 不支持jpg,QPixmap/QImage不能读取jpg图像,但是可以读取bmp,png图像。
		// "bmp", "pbm", "pgm", "png", "ppm", "xbm", "xpm" ,也就是目前只能支持这些格式。
		//var pixmap *qtgui.QPixmap
		//pixmap = pixmap.NewForInherit3p("/images/logo1.png")
		pixmap := qtgui.NewQPixmap3p(":/images/logo1.png")
		mw.Label.SetPixmap(pixmap)
		mw.Label.SetFixedSize1(100, 100)
		mw.Label.SetScaledContents(true)
	}

	if false {
		mw.Label.SetGeometry(10, 10, 200, 200)
		pixmap := qtgui.NewQPixmap3p("://images/logo1.png")
		mw.Label.SetPixmap(pixmap)
		mw.Label.SetFixedSize1(100, 100)
		mw.Label.SetScaledContents(true)
	}

	if false {
		mw.Label.SetGeometry(10, 10, 200, 200)
		pixmap := qtgui.NewQPixmap3p(":/logo1-img.png")
		mw.Label.SetPixmap(pixmap)
		mw.Label.SetFixedSize1(100, 100)
		mw.Label.SetScaledContents(true)
	}

	if false {
		mw.Label.SetGeometry(10, 10, 200, 200)
		pixmap := qtgui.NewQPixmap3p(":/Images/images/g8.png")
		mw.Label.SetPixmap(pixmap)
		mw.Label.SetFixedSize1(100, 100)
		mw.Label.SetScaledContents(true)
	}

	if true { // 05_rcc
		// QLocale__English
		// https://blog.csdn.net/shenenhua/article/details/79150053
		// http://doc.qt.io/archives/qt-4.8/qlocale.html#Language-enum
		// v := qtcore.NewQLocale2p(qtcore.QLocale__English)
		// mw.Label.SetLocale(v)
		// str := v.System().Name()
		// fmt.Println(str)

		mw.Label.SetGeometry(10, 10, 200, 200)
		pixmap := qtgui.NewQPixmap3p(":/Images/g8.png")
		mw.Label.SetPixmap(pixmap)
		mw.Label.SetFixedSize1(100, 100)
		mw.Label.SetScaledContents(true)
		mw.Label.UnsetLocale()
	}

	if false {
		// mw.Label.SetGeometry(10, 10, 200, 200)
		// strText1 := "<html><body> <img src=\":/Images/g8.png\" width=\"100\" height=\"100\" /> </body> </html>"
		// mw.Label.SetText(strText1)
		// mw.Label.SetAlignment(qtcore.Qt__AlignCenter)
		pMovie := qtgui.NewQMovie2p(":/Images/zb.gif")
		mw.Label.SetMovie(pMovie)
		mw.Label.SetFixedSize1(125, 125)
		mw.Label.SetScaledContents(true)
		pMovie.Start()

	}

	mw.Form.Show()
	app.Exec()
}
