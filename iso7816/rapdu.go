package iso7816

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/dibranmulder/gmrtd/utils"
)

const RAPDU_STATUS_SUCCESS = 0x9000
const RAPDU_STATUS_SELECTED_FILE_INVALIDATED = 0x6283 // warning
const RAPDU_SECURITY_CONDITION_NOT_SATIFIED = 0x6982
const RAPDU_STATUS_FUNCTION_NOT_SUPPORTED = 0x6A81
const RAPDU_STATUS_FILENOTFOUND = 0x6A82
const RAPDU_STATUS_INCORRECT_P1_OR_P2_PARAMETER = 0x6A86

type RApdu struct {
	Data   []byte
	Status uint16
}

func NewRApdu(status uint16, data []byte) *RApdu {
	return &RApdu{Status: status, Data: data}
}

func (apdu *RApdu) String() string {
	return fmt.Sprintf("(Status:%04x, Data:%x)", apdu.Status, apdu.Data)
}

func (apdu *RApdu) IsSuccess() bool {
	return apdu.Status == RAPDU_STATUS_SUCCESS
}

func (apdu *RApdu) FileNotFound() bool {
	return apdu.Status == RAPDU_STATUS_FILENOTFOUND
}

func ParseRApdu(data []byte) (rapdu *RApdu, err error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("rApdu length must be >= 2 (Length:%d)", len(data))
	}

	rapdu = new(RApdu)

	rapdu.Status = binary.BigEndian.Uint16(data[len(data)-2:])

	rapdu.Data = make([]byte, len(data)-2)
	copy(rapdu.Data, data[0:len(data)-2])

	return rapdu, nil
}

func (rApdu *RApdu) Encode() []byte {
	var out []byte
	out = bytes.Clone(rApdu.Data)
	out = append(out, utils.UInt16ToBytes(rApdu.Status)...)
	return out
}
