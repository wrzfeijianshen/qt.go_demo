import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0
import Qt.labs.settings 1.0
import "elements"
Item {
    id: window
    width: 360
    height: 520
    visible: true
    Settings {
        id: settings
        property string style: "Default"
    }

    ColumnLayout {
        width: parent.width
        anchors.fill: parent
        ToolBar {
            id: toolbar
            Material.foreground: "white"
            Material.background: Material.BlueGrey
             z: 100
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.top: parent.top
            RowLayout {
                spacing: 20
                anchors.fill: parent

                ToolButton {
                    contentItem: Image {
                        fillMode: Image.Pad
                        horizontalAlignment: Image.AlignHCenter
                        verticalAlignment: Image.AlignVCenter
                        source: "images/drawer.png" //stackView.depth > 1 ? "images/back.png" : "images/drawer.png"
                    }
                    onClicked: {
                        //if (stackView.depth > 1) {
                        //    stackView.pop()
                        //    listView.currentIndex = -1
                        //} else {
                        //    drawer.open()
                        //}
                        drawer.open()
                    }
                }

                Label {
                    id: titleLabel
                    text: listView.currentItem ? listView.currentItem.text : "My App"
                    font.pixelSize: 20
                    elide: Label.ElideRight
                    horizontalAlignment: Qt.AlignHCenter
                    verticalAlignment: Qt.AlignVCenter
                    Layout.fillWidth: true
                }
                ToolButton {
                    id: settingsViewer
                    contentItem: Image {
                        fillMode: Image.Pad
                        horizontalAlignment: Image.AlignHCenter
                        verticalAlignment: Image.AlignVCenter
                        source: "qrc:/qml/images/menu.png"
                    }
                    onClicked: optionsMenu.open()

                    Menu {
                        id: optionsMenu
                        MenuItem {
                            text: "Settings"
                            onTriggered: settingsPopup.open()
                        }
                        MenuItem {
                            text: "About"
                            onTriggered: aboutDialog.open()
                        }
                        MenuItem {
                            text: "Log Out"
                            onTriggered: logoutDialog.open()
                        }
                    }
                }
            }
        }
        Toast {
            //a toast that everyone can use
            id: globalToast
            x: parent.width / 10
            y: (parent.height * 4) / 5
            width: (parent.width * 4) / 5
        }
        BusyIndicator {
            id: loadingIndicator
            visible: true
            z: 100
            readonly property int size: Math.min(footer.availableWidth, footer.availableHeight) / 5
            width: size
            height: size
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.bottom: footer.top
            Material.accent: Material.BlueGrey
        }

        ToolBar {
            id: footer
            Material.foreground: "white"
            Material.background: Material.BlueGrey
             z: 100
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.bottom: parent.bottom
            RowLayout {
                spacing: 20
                anchors.fill: parent

                Label {
                    id: footerLabel
                    text: ""
                    visible: true
                    font.pixelSize: 16
                    elide: Label.ElideRight
                    horizontalAlignment: Qt.AlignHCenter
                    verticalAlignment: Qt.AlignVCenter
                    Layout.fillWidth: true
                }
                ProgressBar {
                    id: progressIndicator
                    value: 0.0
                    indeterminate: false
                    visible: false
                    z: 100
                    width: parent.width
                    anchors.horizontalCenter: parent.horizontalCenter
                    Material.accent: Material.Grey
                }
                Connections {
                    target: QmlBridge
    //Progress bar update
                    onUpdateProcessStatus: {
                        //initialise the viewing
                        footerLabel.visible = false
                        progressIndicator.visible = true
                        progressIndicator.indeterminate = indeterminate
                        //set the progress value (only useful when determinate)
                        progressIndicator.value = c
                        if (c.toFixed(2) ==  1.0) {
                            //process complete
                            progressIndicator.visible = false
                            footerLabel.text = "process complete"
                            footerLabel.visible = true
                        }
                    }
                }
            }
        }
        //content holder
        StackView {
            id: stackView
            Layout.fillWidth: true
            Layout.fillHeight: true
            anchors.margins: 10
              Connections {
                target: QmlBridge
//hotloading:
                onUpdateLoader: {
                    stackView.clear()
                    stackView.push(p)
                    footerLabel.text = "detected change to: " + p
                    loadingIndicator.visible = true
                }
              }

//animates the loader for 1 second when respawning a page for effect
                PropertyAnimation {
                    running: true
                    target: loadingIndicator
                    property: 'visible'
                    to: false
                    duration: 2000
                }

            initialItem: Pane {
                id: pane

                Image {
                    id: logo
                    width: pane.availableWidth / 2
                    height: pane.availableHeight / 2
                    anchors.centerIn: parent
                    anchors.verticalCenterOffset: -50
                    fillMode: Image.PreserveAspectFit
                    source: "qrc:/qml/images/qt-logo.png"
                }

                Label {
                    text: "Qt is a set of controls that can be used to build complete interfaces in Qt Quick."
                    anchors.margins: 20
                    anchors.top: logo.bottom
                    anchors.left: parent.left
                    anchors.right: parent.right
                    horizontalAlignment: Label.AlignHCenter
                    verticalAlignment: Label.AlignVCenter
                    wrapMode: Label.Wrap
                }

            }
            onCurrentItemChanged: currentItem.anchors.fill = this
        }
    }
    //menu
    Drawer {
        id: drawer
        width: Math.min(window.width, window.height) / 3 * 2
        height: window.height
        dragMargin: stackView.depth > 1 ? 0 : undefined

        ListView {
            id: listView
            currentIndex: -1
            anchors.fill: parent

            delegate: ItemDelegate {
                width: parent.width
                text: model.title
                highlighted: ListView.isCurrentItem
                onClicked: {
                    if (listView.currentIndex != index) {
                        listView.currentIndex = index
                        stackView.push(model.source)
                        titleLabel.text = model.title
                    }
                    drawer.close()
                }
            }

            model: ListModel {
            //application
                ListElement { title: "Contacts"; source: "qrc:/qml/pages/_contactsPage.qml" }
                ListElement { title: "Files"; source: "qrc:/qml/pages/_filelistPage.qml" }
                ListElement { title: "Downloads"; source: "qrc:/qml/pages/_downloadsPage.qml" }
                ListElement { title: "Search"; source: "qrc:/qml/pages/_searchPage.qml" }
                ListElement { title: "Login"; source: "qrc:/qml/pages/_loginPage.qml" }
                ListElement { title: "Calulator"; source: "qrc:/qml/pages/_calculatorPage.qml" }
                ListElement { title: "Clock"; source: "qrc:/qml/pages/_clockPage.qml" }
                ListElement { title: "Toast"; source: "qrc:/qml/pages/_toastPage.qml" }
                ListElement { title: "Asynchronous Loading"; source: "qrc:/qml/pages/_asynchronousPage.qml" }
            //demos
                ListElement { title: "BusyIndicator"; source: "qrc:/qml/pages/BusyIndicatorPage.qml" }
                ListElement { title: "Button"; source: "qrc:/qml/pages/ButtonPage.qml" }
                ListElement { title: "CheckBox"; source: "qrc:/qml/pages/CheckBoxPage.qml" }
                ListElement { title: "ComboBox"; source: "qrc:/qml/pages/ComboBoxPage.qml" }
                ListElement { title: "Dial"; source: "qrc:/qml/pages/DialPage.qml" }
                ListElement { title: "Delegates"; source: "qrc:/qml/pages/DelegatePage.qml" }
                ListElement { title: "Drawer"; source: "qrc:/qml/pages/DrawerPage.qml" }
                ListElement { title: "Frame"; source: "qrc:/qml/pages/FramePage.qml" }
                ListElement { title: "GroupBox"; source: "qrc:/qml/pages/GroupBoxPage.qml" }
                ListElement { title: "Menu"; source: "qrc:/qml/pages/MenuPage.qml" }
                ListElement { title: "PageIndicator"; source: "qrc:/qml/pages/PageIndicatorPage.qml" }
                ListElement { title: "Popup"; source: "qrc:/qml/pages/PopupPage.qml" }
                ListElement { title: "ProgressBar"; source: "qrc:/qml/pages/ProgressBarPage.qml" }
                ListElement { title: "RadioButton"; source: "qrc:/qml/pages/RadioButtonPage.qml" }
                ListElement { title: "RangeSlider"; source: "qrc:/qml/pages/RangeSliderPage.qml" }
                ListElement { title: "ScrollBar"; source: "qrc:/qml/pages/ScrollBarPage.qml" }
                ListElement { title: "ScrollIndicator"; source: "qrc:/qml/pages/ScrollIndicatorPage.qml" }
                ListElement { title: "Slider"; source: "qrc:/qml/pages/SliderPage.qml" }
                ListElement { title: "SpinBox"; source: "qrc:/qml/pages/SpinBoxPage.qml" }
                ListElement { title: "StackView"; source: "qrc:/qml/pages/StackViewPage.qml" }
                ListElement { title: "SwipeView"; source: "qrc:/qml/pages/SwipeViewPage.qml" }
                ListElement { title: "Switch"; source: "qrc:/qml/pages/SwitchPage.qml" }
                ListElement { title: "TabBar"; source: "qrc:/qml/pages/TabBarPage.qml" }
                ListElement { title: "TextArea"; source: "qrc:/qml/pages/TextAreaPage.qml" }
                ListElement { title: "TextField"; source: "qrc:/qml/pages/TextFieldPage.qml" }
                ListElement { title: "ToolTip"; source: "qrc:/qml/pages/ToolTipPage.qml" }
                ListElement { title: "Tumbler"; source: "qrc:/qml/pages/TumblerPage.qml" }
            }

            ScrollIndicator.vertical: ScrollIndicator { }
        }
    }
    Popup {
        id: notificationPopup
        x: (window.width - width) / 2
        y: window.height / 6
        width: Math.min(window.width, window.height) / 3 * 2
        height: settingsColumn.implicitHeight + topPadding + bottomPadding
        modal: true
        focus: true
        contentItem: Text {
            id: notifcationText
            text: ""
        }
    }
    Popup {
        id: settingsPopup
        x: (window.width - width) / 2
        y: window.height / 6
        width: Math.min(window.width, window.height) / 3 * 2
        height: settingsColumn.implicitHeight + topPadding + bottomPadding
        modal: true
        focus: true

        contentItem: ColumnLayout {
            id: settingsColumn
            spacing: 20

            Label {
                text: "Configuration"
                font.bold: true
            }

            RowLayout {
                spacing: 10

                GridLayout {
                    id: grid
                    columns: 2

                    Text { text: "Hotloading"; font.bold: true; }
                    Text { id: hotloading; text: "waiting"; color: "red" }
                    Text { text: "Version"; font.bold: true; }
                    Text { id: host; text: "waiting"; color: "red" }
                }

              Connections {
                target: QmlBridge
                onUpdateSettings: {
                    footerLabel.text = version
                    hotloading.text = hotload
                }
              }
                //ComboBox {
                //    id: styleBox
                //    property int styleIndex: -1
                //    model: ["Default", "Material", "Universal"]
                //    Component.onCompleted: {
                //        styleIndex = find(settings.style, Qt.MatchFixedString)
                //        if (styleIndex !== -1)
                //            currentIndex = styleIndex
                //    }
                //    Layout.fillWidth: true
                //}
            }

            //Label {
            //    text: "Restart required"
            //    color: "#e41e25"
            //    opacity: styleBox.currentIndex !== styleBox.styleIndex ? 1.0 : 0.0
            //    horizontalAlignment: Label.AlignHCenter
            //    verticalAlignment: Label.AlignVCenter
            //    Layout.fillWidth: true
            //    Layout.fillHeight: true
            //}

            RowLayout {
                spacing: 10

                Button {
                    id: okButton
                    text: "Ok"
                    onClicked: {
                        //settings.style = styleBox.displayText
                        settingsPopup.close()
                    }

                    Material.foreground: Material.primary
                    Material.background: "transparent"
                    Material.elevation: 0

                    Layout.preferredWidth: 0
                    Layout.fillWidth: true
                }

                Button {
                    id: cancelButton
                    text: "Cancel"
                    onClicked: {
                        //styleBox.currentIndex = styleBox.styleIndex
                        settingsPopup.close()
                    }

                    Material.background: "transparent"
                    Material.elevation: 0

                    Layout.preferredWidth: 0
                    Layout.fillWidth: true
                }
            }
        }
    }
    Popup {
        id: logoutDialog
        modal: true
        focus: true
        x: (window.width - width) / 2
        y: window.height / 6
        width: Math.min(window.width, window.height) / 3 * 2
        contentHeight: logoutColumn.height

        Column {
            id: logoutColumn
            spacing: 20

            Label {
                text: "Logout"
                font.bold: true
            }

            Label {
                width: logoutDialog.availableWidth
                text: "Logging out is really only necessary if you share a computer with others. If you want to continue, click logout below otherwise click elsewhere on the window"
                wrapMode: Label.Wrap
                font.pixelSize: 12
            }
            Button {
                id: searchButton
                text: "Logout"
                width: logoutDialog.availableWidth
                onClicked: function() {
                    console.log('logging out')
                }
            }
        }
    }
    Popup {
        id: aboutDialog
        modal: true
        focus: true
        x: (window.width - width) / 2
        y: window.height / 6
        width: Math.min(window.width, window.height) / 3 * 2
        contentHeight: aboutColumn.height

        Column {
            id: aboutColumn
            spacing: 20

            Label {
                text: "About"
                font.bold: true
            }

            Label {
                width: aboutDialog.availableWidth
                text: "The Qt Quick Controls 2 module delivers the next generation user interface controls based on Qt Quick."
                wrapMode: Label.Wrap
                font.pixelSize: 12
            }

            Label {
                width: aboutDialog.availableWidth
                text: "In comparison to the desktop-oriented Qt Quick Controls 1, Qt Quick Controls 2 "
                    + "are an order of magnitude simpler, lighter and faster, and are primarily targeted "
                    + "towards embedded and mobile platforms."
                wrapMode: Label.Wrap
                font.pixelSize: 12
            }
        }
    }
}
