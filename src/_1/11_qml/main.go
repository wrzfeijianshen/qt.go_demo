package main

import (
	_ "fmt"
	"os"

	"github.com/kitech/qt.go/qtqml"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	qmle := qtqml.NewQQmlApplicationEngine(nil)
	// qmle.AddImportPath(":/main.qml")
	qmle.Load1(":/main.qml")

	app.Exec()
}
