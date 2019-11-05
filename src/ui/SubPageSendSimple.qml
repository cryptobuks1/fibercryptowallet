import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "Controls" // For quick UI development, switch back to resources when making a release

Page {
    id: root
    property string walletSelected
    property bool walletEncrypted: false 
    property string amount
    property string destinationAddress
    function getSelectedWallet(){
        return walletSelected
    }
    function getAmount(){
        return amount
    }
    function getDestinationAddress(){
        return destinationAddress
    }
    function walletIsEncrypted(){
        return walletEncrypted
    }
    signal qrCodeRequested(var data)

    onQrCodeRequested: {
        dialogQR.setVars(data)
        dialogQR.open()
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 10
        anchors.rightMargin: 10
        spacing: 20

        ColumnLayout {
            id: columnLayoutSendFrom

            Layout.alignment: Qt.AlignTop

            Label { text: qsTr("Send from") }

            ComboBox {
                id: comboBoxWalletsSendFrom

                Layout.fillWidth: true
                textRole: "name"
                displayText: comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].sky ? comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].sky + " --- " + comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].name : currentText

                model: WalletModel {
                    Component.onCompleted: {
                        loadModel(walletManager.getWallets())
                    }
                } 
                
                // Taken from Qt 5.13.0 source code:
                delegate: MenuItem {
                    width: parent.width
                    text: comboBoxWalletsSendFrom.textRole ? (Array.isArray(comboBoxWalletsSendFrom.model) ? modelData["sky"] + " --- " + modelData[comboBoxWalletsSendFrom.textRole] : model["sky"] + " --- " + model[comboBoxWalletsSendFrom.textRole]) : " --- " + modelData
                    Material.foreground: comboBoxWalletsSendFrom.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                    highlighted: comboBoxWalletsSendFrom.highlightedIndex === index
                    hoverEnabled: comboBoxWalletsSendFrom.hoverEnabled
                    leftPadding: highlighted ? 2*padding : padding // added
                    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                }

                onActivated: {
                    root.walletSelected = comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].fileName
                    root.walletEncrypted = comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].encryptionEnabled
                }
            } // ComboBox
        } // ColumnLayout (send from)

        ColumnLayout {
            id: columnLayoutSendTo

            Layout.alignment: Qt.AlignTop

            Label { text: qsTr("Send to") }
            
            RowLayout {
                Layout.fillWidth: true
                spacing: 8

                ToolButtonQR {
                    id: toolButtonQR
                    Layout.bottomMargin: 4

                    iconSize: "24x24"

                    onClicked: {
                        qrCodeRequested(textFieldWalletsSendTo.text)
                    }
                }
                
                TextField {
                    id: textFieldWalletsSendTo
                    font.family: "Code New Roman"
                    placeholderText: qsTr("Destination address")
                    selectByMouse: true
                    Layout.fillWidth: true
                    Layout.topMargin: -5
                    onTextChanged:{
                        root.destinationAddress = text
                    }
                }
            } // RowLayout
        } // ColumnLayout (send to)

        TextField {
            id: textFieldAmount
            placeholderText: qsTr("Amount to send")
            selectByMouse: true
            Layout.fillWidth: true
            Layout.topMargin: -10
            validator: DoubleValidator {
                notation: DoubleValidator.StandardNotation
            }
            onTextChanged:{
                    root.amount = text
                }
        }
    } // ColumnLayout (root)
}
