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
func NewUi_Form() *Ui_Form {
	return &Ui_Form{}
}

type Ui_Form struct {
	Form *qtwidgets.QWidget
}

//  struct block end

//  setupUi block begin
// void setupUi(QWidget *Form)
func (this *Ui_Form) SetupUi(Form *qtwidgets.QWidget) {
	this.Form = Form
	// { // 126
	if Form.ObjectName() == "" {
		Form.SetObjectName("Form")
	}
	Form.Resize(400, 300)
	this.Form.SetStyleSheet("background-color: rgb(255, 0, 255);") // 114

	this.RetranslateUi(Form)

	qtcore.QMetaObject_ConnectSlotsByName(Form) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *Form)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_Form) RetranslateUi(Form *qtwidgets.QWidget) {
	// noimpl: {
	this.Form.SetWindowTitle(qtcore.QCoreApplication_Translate("Form", "Form", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_Form2() *Ui_Form {
	this := &Ui_Form{}
	w := qtwidgets.NewQWidget(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_Form) QWidget_PTR() *qtwidgets.QWidget {
	return this.Form.QWidget_PTR()
}

//  done block end
