# 02_画线

http://doc.qt.io/qt-5/qt.html#GlobalColor-enum有些全局的变量如

```
Qt::ImhHiddenText
Qt::Key_Escape
Qt::white
```

在kitech\qt.go\qtcore\qnamespace.go

Qt Namespace 中定义了很多的全局变量值,咱们可以在qnamespace.go 找到对应的名称

如上例子中int 7,就可以Qt__red来代替.

```
		f1 := func(painter *qtgui.QPainter) {
			var qcolor qtgui.QColor

			s := qcolor.NewForInherit1(qtcore.Qt__red)

			painter.SetPen(s)
			painter.DrawLine2(10, 10, 100, 100)
		}
		mw.Form.InheritInitPainter(f1)
```
## Pen Style

