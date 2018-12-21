# 07 Qlabel

添加Widget界面上添加一个Label,go-uic 生成ui.go,在Ui_Form里面有Label *qtwidgets.QLabel控件,咱们就进行操作QLabel,当然一些属性在设计师里面就可以直接修改,添加.

```
type Ui_Form struct {
	Label *qtwidgets.QLabel
	Form  *qtwidgets.QWidget
}
```
我们可以看到 ui.go中的定义构造.
```
this.Label = qtwidgets.NewQLabel(this.Form, 0) // 111
this.Label.SetObjectName("Label")              // 112
this.Label.SetGeometry(40, 30, 54, 12)         // 114

```

通过调用setText可以为标签设置文本（Hello World），这时标签就可以正常显示出来了。为了显示更佳的效果，我们可以通过调用setStyleSheet来设置样式。color: red-顾名思义，就是为标签设置一个文本色（红色）。

在main中

```
	mw := NewUi_Form2()
	mw.Label.SetText("hello world")
	mw.Label.SetGeometry(10, 10, 200, 100)
	mw.Label.SetStyleSheet("color:red")

	mw.Form.Show()
```

对齐方式

默认的标签文本对齐方式为：左对齐、垂直居中，我们可以通过setAlignment来设置，包括：左、上、右、下、居中对齐，一般情况，我们会进行两两组合（水平方向、垂直方向）。
