package document

import (
	"bytes"
	"testing"

	"github.com/dibranmulder/gmrtd/utils"
)

func TestNewEFDIRNoData(t *testing.T) {
	if efDir, err := NewEFDIR(nil); efDir != nil || err != nil {
		t.Errorf("Should be nil when no input data provided")
	}

	if efDir, err := NewEFDIR([]byte{}); efDir != nil || err != nil {
		t.Errorf("Should be nil when no input data provided")
	}
}

func TestNewEfDirBadTlv(t *testing.T) {
	var efDirBytes []byte = utils.HexToBytes("02101234") // invalid TLV encoding - insufficient bytes

	efDir, err := NewEFDIR(efDirBytes)

	if err == nil {
		t.Errorf("Error expected")
	}

	if efDir != nil {
		t.Errorf("EF.DIR not expected for error case")
	}
}

func TestNewEFDIR(t *testing.T) {
	// Table 31. (EF.DIR Format) - 9303 p10
	fileBytes := utils.HexToBytes("61094F07A000000247100161094F07A000000247200161094F07A000000247200261094F07A0000002472003")

	var err error
	var efDir *EFDIR

	efDir, err = NewEFDIR(fileBytes)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(efDir.Application) != 4 {
		t.Errorf("4 entries expected")
	}

	if !bytes.Equal(efDir.Application[0].aid, utils.HexToBytes("A0000002471001")) {
		t.Errorf("Incorrect application #1")
	}

	if !bytes.Equal(efDir.Application[1].aid, utils.HexToBytes("A0000002472001")) {
		t.Errorf("Incorrect application #2")
	}

	if !bytes.Equal(efDir.Application[2].aid, utils.HexToBytes("A0000002472002")) {
		t.Errorf("Incorrect application #3")
	}

	if !bytes.Equal(efDir.Application[3].aid, utils.HexToBytes("A0000002472003")) {
		t.Errorf("Incorrect application #4")
	}
}
