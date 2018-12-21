package main

import (
	"os"

	"github.com/kitech/qt.go/qtcore"

	_ "github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw := NewUi_Form2()
	mw.Label.SetText("hello world")
	mw.Label.SetGeometry(10, 10, 200, 100)
	mw.Label.SetStyleSheet("color:red")

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
	if true {
		strText := "<html><head><style> font{color:red;} #f{font-size:18px; color: green;} </style>  </head>  <body>  <font>I am a</font><font id=\"f\">Qter</font> <br/><br/>    <img src=\"https://csdnimg.cn/pubfooter/images/csdn-kf.png\" width=\"100\" height=\"100\" />      </body>         </html>"

		mw.Label.SetText(strText)
		// 图片不显示.
		// strText1 := "<html><body> <img src=\"https://csdnimg.cn/pubfooter/images/csdn-kf.png\" width=\"100\" height=\"100\" /> </body> </html>"
		// mw.Label.SetText(strText1)
		mw.Label.SetAlignment(qtcore.Qt__AlignCenter)
	}

	mw.Form.Show()
	app.Exec()
}
