package document

import (
	"fmt"
	"slices"

	"github.com/dibranmulder/gmrtd/tlv"
)

const DG15Tag = 0x6F

type DG15 struct {
	RawData                   []byte `json:"rawData,omitempty"`
	SubjectPublicKeyInfoBytes []byte `json:"subjectPublicKeyInfoBytes,omitempty"`
}

func NewDG15(data []byte) (*DG15, error) {
	if len(data) < 1 {
		return nil, nil
	}

	var out *DG15 = new(DG15)

	out.RawData = slices.Clone(data)

	nodes, err := tlv.Decode(out.RawData)
	if err != nil {
		return nil, fmt.Errorf("[NewDG15] error: %w", err)
	}

	rootNode := nodes.GetNode(DG15Tag)

	if !rootNode.IsValidNode() {
		return nil, fmt.Errorf("(NewDG15) root node (%x) missing", DG15Tag)
	}

	out.SubjectPublicKeyInfoBytes = rootNode.GetNode(0x30).Encode()
	if len(out.SubjectPublicKeyInfoBytes) < 1 {
		return nil, fmt.Errorf("(NewDG15) missing SubjectPublicKeyInfo")
	}

	return out, nil
}
