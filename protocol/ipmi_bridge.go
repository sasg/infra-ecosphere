package protocol

// port from OpenIPMI
// Bridge Network Function
const (
	IPMI_CMD_GET_BRIDGE_STATE = 		0x00
	IPMI_CMD_SET_BRIDGE_STATE = 		0x01
	IPMI_CMD_GET_ICMB_ADDRESS = 		0x02
	IPMI_CMD_SET_ICMB_ADDRESS = 		0x03
	IPMI_CMD_SET_BRIDGE_PROXY_ADDRESS = 	0x04
	IPMI_CMD_GET_BRIDGE_STATISTICS = 	0x05
	IPMI_CMD_GET_ICMB_CAPABILITIES = 	0x06

	IPMI_CMD_CLEAR_BRIDGE_STATISTICS = 	0x08
	IPMI_CMD_GET_BRIDGE_PROXY_ADDRESS = 	0x09
	IPMI_CMD_GET_ICMB_CONNECTOR_INFO = 	0x0a
	IPMI_CMD_SET_ICMB_CONNECTOR_INFO = 	0x0b
	IPMI_CMD_SEND_ICMB_CONNECTION_ID = 	0x0c

	IPMI_CMD_PREPARE_FOR_DISCOVERY = 	0x10
	IPMI_CMD_GET_ADDRESSES = 		0x11
	IPMI_CMD_SET_DISCOVERED = 		0x12
	IPMI_CMD_GET_CHASSIS_DEVICE_ID = 	0x13
	IPMI_CMD_SET_CHASSIS_DEVICE_ID = 	0x14

	IPMI_CMD_BRIDGE_REQUEST = 		0x20
	IPMI_CMD_BRIDGE_MESSAGE = 		0x21

	IPMI_CMD_GET_EVENT_COUNT = 		0x30
	IPMI_CMD_SET_EVENT_DESTINATION = 	0x31
	IPMI_CMD_SET_EVENT_RECEPTION_STATE = 	0x32
	IPMI_CMD_SEND_ICMB_EVENT_MESSAGE = 	0x33
	IPMI_CMD_GET_EVENT_DESTIATION = 	0x34
	IPMI_CMD_GET_EVENT_RECEPTION_STATE = 0x35

	IPMI_CMD_ERROR_REPORT = 			0xff
)