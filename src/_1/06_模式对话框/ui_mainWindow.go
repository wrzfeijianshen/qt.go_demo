package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"
import "github.com/kitech/qt.go/qtrt"

func init() { qtcore.KeepMe() }
func init() { qtgui.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }
func init() { qtrt.KeepMe() }

//  header block end

//  struct block begin
func NewUi_MainWindow() *Ui_MainWindow {
	return &Ui_MainWindow{}
}

type Ui_MainWindow struct {
	Centralwidget *qtwidgets.QWidget
	Label         *qtwidgets.QLabel
	Menubar       *qtwidgets.QMenuBar
	Statusbar     *qtwidgets.QStatusBar
	MainWindow    *qtwidgets.QMainWindow
}

//  struct block end

//  setupUi block begin
// void setupUi(QMainWindow *MainWindow)
func (this *Ui_MainWindow) SetupUi(MainWindow *qtwidgets.QMainWindow) {
	this.MainWindow = MainWindow
	// { // 126
	if MainWindow.ObjectName() == "" {
		MainWindow.SetObjectName("MainWindow")
	}
	MainWindow.Resize(368, 292)
	this.Centralwidget = qtwidgets.NewQWidget(this.MainWindow, 0) // 111
	this.Centralwidget.SetObjectName("Centralwidget")             // 112
	this.Label = qtwidgets.NewQLabel(this.Centralwidget, 0)       // 111
	this.Label.SetObjectName("Label")                             // 112
	this.Label.SetGeometry(50, 60, 231, 71)                       // 114
	this.MainWindow.SetCentralWidget(this.Centralwidget)          // 114
	this.Menubar = qtwidgets.NewQMenuBar(this.MainWindow)         // 111
	this.Menubar.SetObjectName("Menubar")                         // 112
	this.Menubar.SetGeometry(0, 0, 368, 23)                       // 114
	this.MainWindow.SetMenuBar(this.Menubar)                      // 114
	this.Statusbar = qtwidgets.NewQStatusBar(this.MainWindow)     // 111
	this.Statusbar.SetObjectName("Statusbar")                     // 112
	this.MainWindow.SetStatusBar(this.Statusbar)                  // 114

	this.RetranslateUi(MainWindow)

	qtcore.QMetaObject_ConnectSlotsByName(MainWindow) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QMainWindow *MainWindow)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MainWindow) RetranslateUi(MainWindow *qtwidgets.QMainWindow) {
	// noimpl: {
	this.MainWindow.SetWindowTitle(qtcore.QCoreApplication_Translate("MainWindow", "MainWindow", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\344\270\273\347\225\214\351\235\242\346\230\276\347\244\272,gogogo", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_MainWindow2() *Ui_MainWindow {
	this := &Ui_MainWindow{}
	w := qtwidgets.NewQMainWindow(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_MainWindow) QWidget_PTR() *qtwidgets.QWidget {
	return this.MainWindow.QWidget_PTR()
}

//  done block end
