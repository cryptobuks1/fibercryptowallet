import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0
import "../"

Item {
    id: root

    readonly property real addressListHeight: listViewFilterAddress.height
    readonly property real delegateHeight: 42
    property alias tristate: checkDelegate.tristate
    property alias walletText: checkDelegate.text
    
    clip: true
    width: 300
    height: checkDelegate.height +  columnLayout.spacing + listViewFilterAddress.height
    
    ColumnLayout {
        id: columnLayout
        anchors.fill: parent

        CheckDelegate {
            id: checkDelegate

            Layout.fillWidth: true
            tristate: true
            text: name
            LayoutMirroring.enabled: true

            nextCheckState: function() {
                if (checkState === Qt.Checked) {
                    if (!listViewFilterAddress.allChecked) {
                        listViewFilterAddress.allChecked = true
                    }
                    listViewFilterAddress.allChecked = false
                    return Qt.Unchecked
                } else {
                    if (listViewFilterAddress.allChecked) {
                        listViewFilterAddress.allChecked = false
                    }
                    listViewFilterAddress.allChecked = true
                    return Qt.Checked
                }
            }

            onCheckStateChanged:{
                if (checkState === Qt.Unchecked){
                    for (var i = 0; i < listViewFilterAddress.listAddresses.addresses.length; i++){
                        var address = listViewFilterAddress.listAddresses.addresses[i]
                        listViewFilterAddress.listAddresses.editAddress(i, address.address, address.sky, address.coinHours, false)
                    }
                } else if(checkState === Qt.Checked){
                    for (var i = 0; i < listViewFilterAddress.listAddresses.addresses.length; i++){
                        var address = listViewFilterAddress.listAddresses.addresses[i]
                        listViewFilterAddress.listAddresses.editAddress(i, address.address, address.sky, address.coinHours, true)
                    }
                }
            }

            contentItem: Label {
                leftPadding: checkDelegate.indicator.width + checkDelegate.spacing
                verticalAlignment: Qt.AlignVCenter
                text: checkDelegate.text
                color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
            }
        } // CheckDelegate

        ListView {
            id: listViewFilterAddress
            property AddressModel listAddresses
            property int checkedDelegates: 0
            property bool allChecked: false
            model: listAddresses
            
            Layout.fillWidth: true
            implicitHeight: contentHeight
            interactive: false

            onCheckedDelegatesChanged: {
                if (checkedDelegates === 0) {
                    checkDelegate.checkState = Qt.Unchecked
                } else if (checkedDelegates === count) {
                    checkDelegate.checkState = Qt.Checked
                } else {
                    checkDelegate.checkState = Qt.PartiallyChecked
                }                
            }
            
            delegate: HistoryFilterListAddressDelegate {
                // BUG: Checking the wallet does not change the check state of addresses
                // Is `checked: marked` ok? Or it should be the opposite?
                checked: marked 
                width: parent.width
                text: address 
                onCheckedChanged: {                   
                    ListView.view.checkedDelegates += checked ? 1: -1
                    
                    if (checked) {
                        historyManager.addFilter(address)
                    } else {
                        historyManager.removeFilter(address)
                    }
                    listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, checked)
                }
            } // HistoryFilterListAddressDelegate
                
            
            
            Component.onCompleted:{
                modelManager.setWalletManager(walletManager)
                listAddresses = modelManager.getAddressModel(fileName)
            }
        } // ListView
    } // ColumnLayout
}
