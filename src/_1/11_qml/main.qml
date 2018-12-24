import QtQuick 2.12
import QtQuick.Window 2.12

Window {
    visible: true
    width: 640
    height: 480
    title: qsTr("Hello World")

    Rectangle {
        x:10
        y:10
        width: 300
        height: 300
        color: "blue"
        gradient: Gradient {
            GradientStop { position: 0.0; color: "yellow" }
            GradientStop { position: 1.0; color: "green" }
        }

        Text {
            anchors.centerIn: parent
            text: "Hello, QML!"
        }
        Text {
            anchors.top: parent
            text: "Hello top!"
        }
    }
    Rectangle {
        x:350
        y:350
        width: 100
        height: 100
        color: "red"

        Text {
            anchors.centerIn: parent
            text: "Hello, QML!"
        }
        Text {
            anchors.top: parent
            text: "Hello top!"
        }
    }
}
