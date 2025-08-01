package iso7816

import (
	"bytes"

	"github.com/dibranmulder/gmrtd/utils"
)

type MockTransceiverReqRsp struct {
	req []byte
	rsp []byte
}

type MockTransceiver struct {
	reqRspArr []MockTransceiverReqRsp
}

func (transceiver *MockTransceiver) AddReqRsp(reqHexStr string, rspHexStr string) {
	item := MockTransceiverReqRsp{}
	item.req = utils.HexToBytes(reqHexStr)
	item.rsp = utils.HexToBytes(rspHexStr)

	transceiver.reqRspArr = append(transceiver.reqRspArr, item)
}

func (transceiver *MockTransceiver) Transceive(cla int, ins int, p1 int, p2 int, data []byte, le int, encodedData []byte) []byte {
	// NB we ignore the raw cApdu fields (cla,ins,p1,p2,data,l2) and just use encodedData (which is the same)
	for i := range transceiver.reqRspArr {
		if bytes.Equal(transceiver.reqRspArr[i].req, encodedData) {
			return bytes.Clone(transceiver.reqRspArr[i].rsp)
		}
	}

	// if we got here then we couldn't match the C-APDU
	return []byte{}
}
