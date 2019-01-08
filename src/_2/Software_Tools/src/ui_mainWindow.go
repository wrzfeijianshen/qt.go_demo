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
	Actiond         *qtwidgets.QAction
	Actions         *qtwidgets.QAction
	Actiona         *qtwidgets.QAction
	Centralwidget   *qtwidgets.QWidget
	GridLayout      *qtwidgets.QGridLayout
	TabWidget       *qtwidgets.QTabWidget
	Tab_3           *qtwidgets.QWidget
	GroupBox        *qtwidgets.QGroupBox
	Label           *qtwidgets.QLabel
	Label_2         *qtwidgets.QLabel
	LineEdit        *qtwidgets.QLineEdit
	BtnOpen         *qtwidgets.QPushButton
	BtnReadHosts    *qtwidgets.QPushButton
	BtnSaveHosts    *qtwidgets.QPushButton
	PlainTextEdit   *qtwidgets.QPlainTextEdit
	Tab             *qtwidgets.QWidget
	Tab_4           *qtwidgets.QWidget
	GroupBox_2      *qtwidgets.QGroupBox
	LineEdit_2      *qtwidgets.QLineEdit
	Label_3         *qtwidgets.QLabel
	PushButton      *qtwidgets.QPushButton
	PlainTextEdit_2 *qtwidgets.QPlainTextEdit
	ComboBox        *qtwidgets.QComboBox
	PushButton_2    *qtwidgets.QPushButton
	PushButton_3    *qtwidgets.QPushButton
	LineEdit_3      *qtwidgets.QLineEdit
	Label_4         *qtwidgets.QLabel
	PushButton_4    *qtwidgets.QPushButton
	CheckBox        *qtwidgets.QCheckBox
	Tab_2           *qtwidgets.QWidget
	Menubar         *qtwidgets.QMenuBar
	Menu            *qtwidgets.QMenu
	Menu_2          *qtwidgets.QMenu
	Menu_3          *qtwidgets.QMenu
	Menu_4          *qtwidgets.QMenu
	Statusbar       *qtwidgets.QStatusBar
	MainWindow      *qtwidgets.QMainWindow
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
	MainWindow.Resize(500, 500)
	this.Actiond = qtwidgets.NewQAction(this.MainWindow)                // 111
	this.Actiond.SetObjectName("Actiond")                               // 112
	this.Actions = qtwidgets.NewQAction(this.MainWindow)                // 111
	this.Actions.SetObjectName("Actions")                               // 112
	this.Actiona = qtwidgets.NewQAction(this.MainWindow)                // 111
	this.Actiona.SetObjectName("Actiona")                               // 112
	this.Centralwidget = qtwidgets.NewQWidget(this.MainWindow, 0)       // 111
	this.Centralwidget.SetObjectName("Centralwidget")                   // 112
	this.GridLayout = qtwidgets.NewQGridLayout(this.Centralwidget)      // 111
	this.GridLayout.SetObjectName("GridLayout")                         // 112
	this.TabWidget = qtwidgets.NewQTabWidget(this.Centralwidget)        // 111
	this.TabWidget.SetObjectName("TabWidget")                           // 112
	this.Tab_3 = qtwidgets.NewQWidget(nil, 0)                           // 111
	this.Tab_3.SetObjectName("Tab_3")                                   // 112
	this.GroupBox = qtwidgets.NewQGroupBox(this.Tab_3)                  // 111
	this.GroupBox.SetObjectName("GroupBox")                             // 112
	this.GroupBox.SetGeometry(10, 10, 461, 311)                         // 114
	this.Label = qtwidgets.NewQLabel(this.GroupBox, 0)                  // 111
	this.Label.SetObjectName("Label")                                   // 112
	this.Label.SetGeometry(20, 30, 51, 21)                              // 114
	this.Label_2 = qtwidgets.NewQLabel(this.GroupBox, 0)                // 111
	this.Label_2.SetObjectName("Label_2")                               // 112
	this.Label_2.SetGeometry(20, 70, 101, 21)                           // 114
	this.LineEdit = qtwidgets.NewQLineEdit(this.GroupBox)               // 111
	this.LineEdit.SetObjectName("LineEdit")                             // 112
	this.LineEdit.SetGeometry(110, 70, 231, 20)                         // 114
	this.BtnOpen = qtwidgets.NewQPushButton(this.GroupBox)              // 111
	this.BtnOpen.SetObjectName("BtnOpen")                               // 112
	this.BtnOpen.SetGeometry(370, 70, 75, 23)                           // 114
	this.BtnReadHosts = qtwidgets.NewQPushButton(this.GroupBox)         // 111
	this.BtnReadHosts.SetObjectName("BtnReadHosts")                     // 112
	this.BtnReadHosts.SetGeometry(20, 110, 75, 23)                      // 114
	this.BtnSaveHosts = qtwidgets.NewQPushButton(this.GroupBox)         // 111
	this.BtnSaveHosts.SetObjectName("BtnSaveHosts")                     // 112
	this.BtnSaveHosts.SetGeometry(20, 150, 75, 23)                      // 114
	this.PlainTextEdit = qtwidgets.NewQPlainTextEdit(this.GroupBox)     // 111
	this.PlainTextEdit.SetObjectName("PlainTextEdit")                   // 112
	this.PlainTextEdit.SetGeometry(110, 110, 341, 191)                  // 114
	this.TabWidget.AddTab(this.Tab_3, "")                               // 115
	this.Tab = qtwidgets.NewQWidget(nil, 0)                             // 111
	this.Tab.SetObjectName("Tab")                                       // 112
	this.TabWidget.AddTab(this.Tab, "")                                 // 115
	this.Tab_4 = qtwidgets.NewQWidget(nil, 0)                           // 111
	this.Tab_4.SetObjectName("Tab_4")                                   // 112
	this.GroupBox_2 = qtwidgets.NewQGroupBox(this.Tab_4)                // 111
	this.GroupBox_2.SetObjectName("GroupBox_2")                         // 112
	this.GroupBox_2.SetGeometry(10, 10, 461, 391)                       // 114
	this.LineEdit_2 = qtwidgets.NewQLineEdit(this.GroupBox_2)           // 111
	this.LineEdit_2.SetObjectName("LineEdit_2")                         // 112
	this.LineEdit_2.SetGeometry(70, 30, 281, 20)                        // 114
	this.Label_3 = qtwidgets.NewQLabel(this.GroupBox_2, 0)              // 111
	this.Label_3.SetObjectName("Label_3")                               // 112
	this.Label_3.SetGeometry(10, 30, 54, 21)                            // 114
	this.PushButton = qtwidgets.NewQPushButton(this.GroupBox_2)         // 111
	this.PushButton.SetObjectName("PushButton")                         // 112
	this.PushButton.SetGeometry(370, 30, 61, 23)                        // 114
	this.PlainTextEdit_2 = qtwidgets.NewQPlainTextEdit(this.GroupBox_2) // 111
	this.PlainTextEdit_2.SetObjectName("PlainTextEdit_2")               // 112
	this.PlainTextEdit_2.SetGeometry(10, 120, 251, 261)                 // 114
	this.ComboBox = qtwidgets.NewQComboBox(this.GroupBox_2)             // 111
	this.ComboBox.AddItem("", qtcore.NewQVariant12("wtf"))              // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant12("wtf"))              // 115
	this.ComboBox.SetObjectName("ComboBox")                             // 112
	this.ComboBox.SetGeometry(270, 120, 101, 22)                        // 114
	this.PushButton_2 = qtwidgets.NewQPushButton(this.GroupBox_2)       // 111
	this.PushButton_2.SetObjectName("PushButton_2")                     // 112
	this.PushButton_2.SetGeometry(380, 120, 61, 23)                     // 114
	this.PushButton_3 = qtwidgets.NewQPushButton(this.GroupBox_2)       // 111
	this.PushButton_3.SetObjectName("PushButton_3")                     // 112
	this.PushButton_3.SetGeometry(210, 60, 75, 23)                      // 114
	this.LineEdit_3 = qtwidgets.NewQLineEdit(this.GroupBox_2)           // 111
	this.LineEdit_3.SetObjectName("LineEdit_3")                         // 112
	this.LineEdit_3.SetGeometry(65, 62, 133, 20)                        // 114
	this.Label_4 = qtwidgets.NewQLabel(this.GroupBox_2, 0)              // 111
	this.Label_4.SetObjectName("Label_4")                               // 112
	this.Label_4.SetGeometry(11, 61, 48, 16)                            // 114
	this.PushButton_4 = qtwidgets.NewQPushButton(this.GroupBox_2)       // 111
	this.PushButton_4.SetObjectName("PushButton_4")                     // 112
	this.PushButton_4.SetGeometry(290, 60, 75, 23)                      // 114
	this.CheckBox = qtwidgets.NewQCheckBox(this.GroupBox_2)             // 111
	this.CheckBox.SetObjectName("CheckBox")                             // 112
	this.CheckBox.SetGeometry(70, 90, 91, 21)                           // 114
	this.CheckBox.SetChecked(true)                                      // 114
	this.TabWidget.AddTab(this.Tab_4, "")                               // 115
	this.Tab_2 = qtwidgets.NewQWidget(nil, 0)                           // 111
	this.Tab_2.SetObjectName("Tab_2")                                   // 112
	this.TabWidget.AddTab(this.Tab_2, "")                               // 115

	this.GridLayout.AddWidget2(this.TabWidget, 0, 0, 1, 1, 0) // 115

	this.MainWindow.SetCentralWidget(this.Centralwidget)      // 114
	this.Menubar = qtwidgets.NewQMenuBar(this.MainWindow)     // 111
	this.Menubar.SetObjectName("Menubar")                     // 112
	this.Menubar.SetGeometry(0, 0, 500, 23)                   // 114
	this.Menu = qtwidgets.NewQMenu(this.Menubar)              // 111
	this.Menu.SetObjectName("Menu")                           // 112
	this.Menu_2 = qtwidgets.NewQMenu(this.Menubar)            // 111
	this.Menu_2.SetObjectName("Menu_2")                       // 112
	this.Menu_3 = qtwidgets.NewQMenu(this.Menubar)            // 111
	this.Menu_3.SetObjectName("Menu_3")                       // 112
	this.Menu_4 = qtwidgets.NewQMenu(this.Menubar)            // 111
	this.Menu_4.SetObjectName("Menu_4")                       // 112
	this.MainWindow.SetMenuBar(this.Menubar)                  // 114
	this.Statusbar = qtwidgets.NewQStatusBar(this.MainWindow) // 111
	this.Statusbar.SetObjectName("Statusbar")                 // 112
	this.MainWindow.SetStatusBar(this.Statusbar)              // 114

	this.Menubar.QWidget.AddAction(this.Menu.MenuAction())   // 115
	this.Menubar.QWidget.AddAction(this.Menu_2.MenuAction()) // 115
	this.Menubar.QWidget.AddAction(this.Menu_3.MenuAction()) // 115
	this.Menubar.QWidget.AddAction(this.Menu_4.MenuAction()) // 115
	this.Menu.QWidget.AddAction(this.Actiond)                // 115
	this.Menu.QWidget.AddAction(this.Actions)                // 115
	this.Menu_4.QWidget.AddAction(this.Actiona)              // 115

	this.RetranslateUi(MainWindow)

	this.TabWidget.SetCurrentIndex(2) // 114

	qtcore.QMetaObject_ConnectSlotsByName(MainWindow) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QMainWindow *MainWindow)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MainWindow) RetranslateUi(MainWindow *qtwidgets.QMainWindow) {
	// noimpl: {
	this.MainWindow.SetWindowTitle(qtcore.QCoreApplication_Translate("MainWindow", "MainWindow", "dummy123", 0))
	this.Actiond.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\346\211\223\345\274\200", "dummy123", 0))
	this.Actions.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\351\200\200\345\207\272", "dummy123", 0))
	this.Actiona.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\345\205\263\344\272\216", "dummy123", 0))
	this.GroupBox.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "hosts\346\226\207\344\273\266", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("MainWindow", "windows:", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "hosts\346\226\207\344\273\266\350\267\257\345\276\204", "dummy123", 0))
	this.LineEdit.SetText(qtcore.QCoreApplication_Translate("MainWindow", "C:\\Windows\\System32\\drivers\\etc", "dummy123", 0))
	this.BtnOpen.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\346\211\223\345\274\200\346\226\207\344\273\266\345\244\271", "dummy123", 0))
	this.BtnReadHosts.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\350\257\273\345\217\226hosts", "dummy123", 0))
	this.BtnSaveHosts.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\344\277\235\345\255\230hosts", "dummy123", 0))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Tab_3), qtcore.QCoreApplication_Translate("MainWindow", "\347\263\273\347\273\237\345\267\245\345\205\267", "dummy123", 0))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Tab), qtcore.QCoreApplication_Translate("MainWindow", "\344\277\256\346\224\271\346\227\266\351\227\264", "dummy123", 0))
	this.GroupBox_2.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "\347\274\226\347\240\201\350\275\254\346\215\242", "dummy123", 0))
	this.LineEdit_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "C:/test/", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\346\226\207\344\273\266\350\267\257\345\276\204", "dummy123", 0))
	this.PushButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\346\265\217\350\247\210", "dummy123", 0))
	this.ComboBox.SetItemText(0, qtcore.QCoreApplication_Translate("MainWindow", "UTF-8", "dummy123", 0))
	this.ComboBox.SetItemText(1, qtcore.QCoreApplication_Translate("MainWindow", "utf-16", "dummy123", 0))
	// noimpl:
	this.PushButton_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\350\275\254\346\215\242\347\274\226\347\240\201", "dummy123", 0))
	this.PushButton_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\346\237\245\346\211\276", "dummy123", 0))
	this.LineEdit_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "*.*", "dummy123", 0))
	this.Label_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\346\226\207\344\273\266\347\261\273\345\236\213", "dummy123", 0))
	this.PushButton_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\350\257\273\345\217\226\347\274\226\347\240\201", "dummy123", 0))
	this.CheckBox.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\345\214\205\345\220\253\345\255\220\347\233\256\345\275\225", "dummy123", 0))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Tab_4), qtcore.QCoreApplication_Translate("MainWindow", "\347\274\226\347\240\201\350\275\254\346\215\242", "dummy123", 0))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Tab_2), qtcore.QCoreApplication_Translate("MainWindow", "\344\270\262\345\217\243\345\267\245\345\205\267", "dummy123", 0))
	this.Menu.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "\346\226\207\344\273\266", "dummy123", 0))
	this.Menu_2.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "\347\274\226\350\276\221", "dummy123", 0))
	this.Menu_3.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "\350\256\276\347\275\256", "dummy123", 0))
	this.Menu_4.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "\345\270\256\345\212\251", "dummy123", 0))
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
