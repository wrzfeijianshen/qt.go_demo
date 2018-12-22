package main

import (
	"os"

	"github.com/kitech/qt.go/qtgui"

	"github.com/kitech/qt.go/qtcore"

	_ "github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
	if false {
		mw.Label.SetText("hello world")
		mw.Label.SetGeometry(10, 10, 200, 100)
		mw.Label.SetStyleSheet("color:red")
	}

	if false {
		mw.Label.SetAlignment(qtcore.Qt__AlignCenter) // 居中对齐
		// 使用样式表来控制（水平居右、垂直居下）：
		//mw.Label.SetStyleSheet("qproperty-alignment: 'AlignBottom | AlignRight';")

		// 自动换行
		mw.Label.SetWordWrap(true)
		strText := "一去二三里,烟村四五家.亭台六七座,八九十枝花."

		mw.Label.SetText(strText)

	}

	if false {
		// 设置行高
		// ### QString
		// Except `QString` class wrapper itself, all other `QString` are replaced with go `string`.
		// 故可以按照string进行操作.如下
		strText := "<p style='line-height:150%'>"
		strText += "一去二三里,烟村四五家.亭台六七座,八九十枝花."
		strText += "</p>"
		// 自动换行
		mw.Label.SetWordWrap(true)

		mw.Label.SetText(strText)
	}

	// 省略
	if false {
		strText := "一去二三里,烟村四五家.亭台六七座,八九十枝花."
		strElidedText := mw.Label.FontMetrics().ElidedText(strText, qtcore.Qt__ElideRight, 100, qtcore.Qt__TextShowMnemonic)
		mw.Label.SetText(strElidedText)
	}

	// 垂直显示
	if false {
		strText := "一去二三里,烟村四五家.亭台六七座,八九十枝花."
		// 思路 每个字添加\n,自己写方法实现哦
		strText = "一\n去\n二\n三\n里\n,\n烟\n村\n四\n五\n家\n.\n亭\n台\n六\n七\n座\n,\n八\n九\n十\n枝\n花\n.\n"
		mw.Label.SetText(strText)
		mw.Label.SetAlignment(qtcore.Qt__AlignCenter)
	}

	// 富文本
	if false {
		strText := "<html><head><style> font{color:red;} #f{font-size:18px; color: green;} </style>  </head>  <body>  <font>I am a</font><font id=\"f\">Qter</font> <br/><br/>    <img src=\"https://csdnimg.cn/pubfooter/images/csdn-kf.png\" width=\"100\" height=\"100\" />      </body>         </html>"

		mw.Label.SetText(strText)
		// 图片不显示,待研究
		strText1 := "<html><body> <img src=\"https://csdnimg.cn/pubfooter/images/csdn-kf.png\" width=\"100\" height=\"100\" /> </body> </html>"
		mw.Label.SetText(strText1)
		mw.Label.SetAlignment(qtcore.Qt__AlignCenter)
	}

	if false {
		var pixmap *qtgui.QPixmap
		// 不支持jpg,QPixmap/QImage不能读取jpg图像,但是可以读取bmp,png图像。
		// "bmp", "pbm", "pgm", "png", "ppm", "xbm", "xpm" ,也就是目前只能支持这些格式。
		pixmap = pixmap.NewForInherit3p("./Images/logo2.bmp")
		mw.Label.SetPixmap(pixmap)
		mw.Label.SetFixedSize1(100, 100)
		mw.Label.SetScaledContents(true)
	}

	if false {
		// 好像不支持jpg动画
		var pMovie *qtgui.QMovie
		pMovie = pMovie.NewForInherit2p("d:\\zb.gif")
		mw.Label.SetMovie(pMovie)
		mw.Label.SetFixedSize1(125, 125)
		mw.Label.SetScaledContents(true)
		pMovie.Start()
	}

	if false {
		// mw.Label.SetNum(666)
		mw.Label.SetNum1(66.666) // 显示数字
	}

	// 超链接
	if true {
		mw.Label.SetText("<a href = \"http://blog.csdn.net/wrzfeijianshen\">飞剑神</a>")
		mw.Label.SetOpenExternalLinks(true)
	}
	mw.Form.Show()
	app.Exec()
}
